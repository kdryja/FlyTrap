package flytrap

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
	"golang.org/x/sync/errgroup"
)

type Proxy struct {
	dst    string
	lc, rc net.Conn
	a      *Auth
}

var (
	connackDenied = []byte{0x20, 0x03, 0x00, 0x87, 0x00}
	subackDenied  = []byte{0x90, 0x04, 0x00, 0x01, 0x00, 0x87}
	pubackDenied  = []byte{0x40, 0x04, 0x00, 0x01, 0x87, 0x00}
)

func readFull(r io.ReadCloser) ([]byte, error) {
	// FUp to the first five bytes of the packet represent Control type and remaining length - needed to read the entire packet.
	h := make([]byte, 1)
	_, err := io.ReadFull(r, h)
	if err != nil {
		return nil, err
	}

	var varHeader []byte

	multiplier := 1
	var value int
	for {
		eB := make([]byte, 1)
		_, err := io.ReadFull(r, eB)
		if err != nil {
			return nil, err
		}
		varHeader = append(varHeader, eB...)
		b := int(eB[0])
		value += (b & 127) * multiplier
		if multiplier > 128*128*128 {
			return nil, fmt.Errorf("malformed variable byte integer")
		}
		multiplier *= 128
		if b&128 == 0 {
			break
		}
	}
	log.Print(value)

	v := make([]byte, value)
	_, err = io.ReadFull(r, v)
	if err != nil {
		return nil, err
	}
	return append(append(h, varHeader...), v...), nil
}

func (p *Proxy) proxyPass(ctx context.Context, lc, rc net.Conn) error {
	var pubKey common.Address
	subPerms := make(map[string]struct{})
	pubPerms := make(map[string]struct{})
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		data, err := readFull(lc)
		if err != nil {
			return err
		}
		log.Print(hex.Dump(data))
		var conErr error
		// Check if it's CONNECT packet
		if data[0] == 0x10 {
			sep := []byte(PROP)
			// Locate flytrap specific property in the payload.
			i := bytes.Index(data, sep)
			// -1 means that token was not found
			if i == -1 {
				// Write CONNACK packet to client, denying access
				if _, err := lc.Write(connackDenied); err != nil {
					conErr = err
				}
				rc.Close()
				return fmt.Errorf("flytrap flag not provided, access denied")
			}
			// Location of the secret key / value property.
			loc := i + len(sep)
			keyLen := binary.BigEndian.Uint16(data[loc : loc+2])
			log.Print(keyLen)
			sigBytes := make([]byte, keyLen)
			copy(sigBytes, data[loc+2:loc+2+int(keyLen)])
			// Attempting to split the payload into signature and compressed public key
			received := bytes.SplitN(sigBytes, []byte{0x00, 0x00}, 2)
			log.Printf("%s", sigBytes)
			sig := received[0]
			pubBytes, err := crypto.DecompressPubkey(received[1])
			if err != nil {
				conErr = err
			}
			sigPub := crypto.PubkeyToAddress(*pubBytes)
			pub, err := crypto.SigToPub(sigPub.Hash().Bytes(), sig)
			if err != nil {
				conErr = err
			}
			if crypto.PubkeyToAddress(*pub) != sigPub {
				conErr = fmt.Errorf("signature was forged")
			}
			log.Printf("Provided pubkey: %s", sigPub.String())
			pubKey = sigPub
		}
		// Verifying provided pubkey failed, cancelling proxy
		if conErr != nil {
			lc.Write(connackDenied)
			rc.Close()
			return conErr
		}
		// Check if it's SUBSCRIBE packet
		if data[0]&0xf0 == 0x80 {
			topics := []string{}
			propertyLen := binary.BigEndian.Uint16(data[2:4])
			start := propertyLen + 4
			for {
				topicLen := binary.BigEndian.Uint16(data[start : start+2])
				topics = append(topics, string(data[start+2:start+2+topicLen]))
				start = start + 3 + topicLen
				if int(start) == len(data) {
					break
				}
			}
			for _, topic := range topics {
				// Check if permission was already verified, if so, skip blockchain communication, as it's costly
				if _, ok := subPerms[topic]; ok {
					log.Printf("Client %q was authorised to subscribe to topic %q", pubKey.String(), topic)
					continue
				}
				ok, err := blockchain.VerifyAccess(topic, pubKey, false)
				if err != nil {
					return err
				}
				// Check if pubkey is authorized to sub to given topic
				if ok {
					// Cache the result for future requests
					subPerms[topic] = struct{}{}
					log.Printf("Client %q was authorised to subscribe to topic %q", pubKey.String(), topic)
					continue
				}
				log.Printf("Client %q attempted subscribe to topic %q and was denied due to insufficient permission", pubKey.String(), topic)
				// Write SUBACK packet to client, denying access
				if _, err := lc.Write(subackDenied); err != nil {
					return err
				}
			}
		}
		// Check if it's PUBLISH packet, coming from the client
		if data[0]&0xf0 == 0x30 && p.lc == lc {
			topLen := data[3+data[2]]
			topic := string(data[4+data[2] : 4+data[2]+topLen])
			// Check if permission was already verified, if so, skip blockchain communication, as it's costly
			if _, ok := pubPerms[topic]; !ok {
				ok, err := blockchain.VerifyAccess(topic, pubKey, true)
				if err != nil {
					return err
				}
				// Check if pubkey is authorized to publish to given topic
				if !ok {
					log.Printf("Client %q attempted publish to topic %q and was denied due to insufficient permission", pubKey.String(), topic)
					// Write PUBACK packet to client, denying access
					if _, err := lc.Write(pubackDenied); err != nil {
						return err
					}
				}
			}
			// Cache the result for future requests
			pubPerms[topic] = struct{}{}
			log.Printf("Client %q was authorised to publish to topic %q", pubKey.String(), topic)
		}
		log.Print("writing response now")
		if _, err := rc.Write(data); err != nil {
			return err
		}
	}
}

// New creates a new instance of Proxy, which is either TLS encrypted or not.
func New(dst string, c net.Conn, s bool, a *Auth) (*Proxy, error) {
	p := &Proxy{dst: dst, lc: c, a: a}
	var err error
	if s {
		rootCA, err := ioutil.ReadFile("mosquitto.org.crt")
		if err != nil {
			return nil, err
		}
		crt := x509.NewCertPool()
		crt.AppendCertsFromPEM(rootCA)
		p.rc, err = tls.Dial("tcp", dst, &tls.Config{RootCAs: crt})
		if err != nil {
			return nil, err
		}
		return p, nil
	}
	p.rc, err = net.Dial("tcp", dst)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Handle starts a bidirectional proxy between provided connection and destination.
func (p *Proxy) Handle() {
	log.Printf("Starting new proxy connection (%s >> %s)\n", p.lc.RemoteAddr().String(), p.dst)
	defer p.rc.Close()
	defer p.lc.Close()
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return p.proxyPass(ctx, p.lc, p.rc)
	})
	g.Go(func() error {
		return p.proxyPass(ctx, p.rc, p.lc)
	})
	if err := g.Wait(); err != nil && err != io.EOF {
		log.Println(err)
		return
	}
	log.Printf("Proxy Connection terminated gracefully. (%s >> %s)", p.lc.RemoteAddr().String(), p.dst)
}
