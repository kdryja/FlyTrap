package flytrap

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/kdryja/thesis/code/flytrap/blockchain"
	"golang.org/x/sync/errgroup"
)

type Proxy struct {
	dst    string
	lc, rc net.Conn
	a      *Auth
}

var subackDenied = []byte{0x90, 0x04, 0x00, 0x01, 0x00, 0x87}

func readFull(r io.ReadCloser) ([]byte, error) {
	// First two bytes of the packet represent Control type and remaining length - needed to read the entire packet.
	h := make([]byte, 2)
	_, err := io.ReadFull(r, h)
	if err != nil {
		return nil, err
	}
	// Second byte indicates remaining length of the packet.
	v := make([]byte, h[1])
	_, err = io.ReadFull(r, v)
	if err != nil {
		return nil, err
	}
	return append(h, v...), nil
}

func proxyPass(ctx context.Context, lc, rc net.Conn, p *Proxy) error {
	var publicKey string
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
		// Check if it's CONNECT packet
		if data[0] == 0x10 {
			sep := []byte(PROP)
			// Locate flytrap specific property in the payload.
			i := bytes.Index(data, sep)
			// -1 means that token was not found
			if i == -1 {
				// Write SUBACK packet to client, denying access
				if _, err := lc.Write(subackDenied); err != nil {
					return err
				}
			}
			// Location of the secret key / value property.
			loc := i + len(sep)
			keyLen := binary.BigEndian.Uint16(data[loc : loc+2])
			publicKey = string(data[loc+2 : loc+2+int(keyLen)])
		}
		// Check if it's SUBSCRIBE packet
		if data[0] == 0x82 {
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
				if _, ok := pubPerms[topic]; ok {
					continue
				}
				ok, err := blockchain.Verify(topic, publicKey, true)
				if err != nil {
					return err
				}
				// Check if pubkey is placed on blockchain
				if ok {
					// Cache the result for future requests
					pubPerms[topic] = struct{}{}
					continue
				}
				// Write SUBACK packet to client, denying access
				if _, err := lc.Write(subackDenied); err != nil {
					return err
				}
			}
		}

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
		return proxyPass(ctx, p.lc, p.rc, p)
	})
	g.Go(func() error {
		return proxyPass(ctx, p.rc, p.lc, p)
	})
	if err := g.Wait(); err != nil && err != io.EOF {
		log.Println(err)
		return
	}
	log.Printf("Proxy Connection terminated gracefully. (%s >> %s)", p.lc.RemoteAddr().String(), p.dst)
}