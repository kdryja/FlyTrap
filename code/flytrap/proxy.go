package flytrap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/eclipse/paho.golang/packets"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
	"golang.org/x/sync/errgroup"
)

type Cache struct {
	Perms *sync.Map
}

type Proxy struct {
	dst    string
	lc, rc net.Conn
	cache  *Cache
}

var (
	connackDenied = []byte{0x20, 0x03, 0x00, 0x87, 0x00}
	subackDenied  = []byte{0x90, 0x04, 0x00, 0x01, 0x00, 0x87}
	pubackDenied  = []byte{0x40, 0x04, 0x00, 0x01, 0x87, 0x00}
)

func (p *Proxy) verifyAccess(pubKey common.Address, topic string, mask byte) (bool, byte, *sync.Map, error) {
	v, _ := p.cache.Perms.LoadOrStore(topic, new(sync.Map))
	permMap, ok := v.(*sync.Map)
	if !ok {
		return false, 0, nil, fmt.Errorf("sync.Map not found in cache")
	}
	result, ok := permMap.LoadOrStore(pubKey, byte(0))
	resultByte, ok := result.(byte)
	if !ok {
		return false, 0, nil, fmt.Errorf("non-byte value found in permission map")
	}
	return ok && resultByte&mask == mask, resultByte, permMap, nil
}

func (p *Proxy) proxyPass(ctx context.Context, lc, rc net.Conn) error {
	defer ctx.Done()
	var pubKey common.Address
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		// For messages received from broker, simply send them back to client
		recv, err := packets.ReadPacket(lc)
		if err != nil {
			return err
		}
		switch {
		// Check if it's CONNECT packet
		case recv.Type == packets.CONNECT:
			conPack := recv.Content.(*packets.Connect)
			sigEncoded, ok := conPack.Properties.User["flytrap_sig"]
			if !ok {
				return fmt.Errorf("flytrap signature not provided")
			}
			sig, err := base64.StdEncoding.DecodeString(sigEncoded)
			if err != nil {
				return err
			}
			pubCompressedEncoded, ok := conPack.Properties.User["flytrap_pub"]
			if !ok {
				return fmt.Errorf("flytrap public key not provided")
			}
			pubCompressed, err := base64.StdEncoding.DecodeString(pubCompressedEncoded)
			if err != nil {
				return err
			}
			pub, err := crypto.DecompressPubkey(pubCompressed)
			if err != nil {
				return err
			}
			if ok := crypto.VerifySignature(crypto.FromECDSAPub(pub), crypto.PubkeyToAddress(*pub).Hash().Bytes(), sig[:len(sig)-1]); !ok {
				return fmt.Errorf("public key signature check failed")
			}
			pubKey = crypto.PubkeyToAddress(*pub)
		// Check if it's SUBSCRIBE packet
		case recv.Type == packets.SUBSCRIBE:
			subPack := recv.Content.(*packets.Subscribe)
			for topic, _ := range subPack.Subscriptions {
				cached, resultByte, permMap, err := p.verifyAccess(pubKey, topic, 2)
				if err != nil {
					return err
				}
				if cached {
					log.Printf("Client %q was already authorised to subscribe to topic %q", pubKey.String(), topic)
					continue
				}
				ok, err := blockchain.VerifyAccess(topic, pubKey, false)
				if err != nil {
					return err
				}
				// Check with blockchain if pubkey is authorized to sub to given topic
				if ok {
					permMap.Store(pubKey, resultByte|2)
					log.Printf("Client %q was authorised to subscribe to topic %q", pubKey.String(), topic)
					continue
				}

				log.Printf("Client %q attempted subscribe to topic %q and was denied due to insufficient permission", pubKey.String(), topic)
				// Write SUBACK packet to client, denying access
				if _, err := lc.Write(subackDenied); err != nil {
					return err
				}
			}
		// Check if it's PUBLISH packet
		case recv.Type == packets.PUBLISH && p.lc == lc:
			pubPack := recv.Content.(*packets.Publish)
			topic := pubPack.Topic
			// Check if permission was already verified, if so, skip blockchain communication, as it's costly
			cached, resultByte, permMap, err := p.verifyAccess(pubKey, topic, 1)
			if err != nil {
				return err
			}
			if cached {
				log.Printf("Client %q was already authorised to publish to topic %q", pubKey.String(), topic)
				break
			}
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
			permMap.Store(pubKey, resultByte|1)
			// Cache the result for future requests
			log.Printf("Client %q was authorised to publish to topic %q", pubKey.String(), topic)
		}
		if _, err := recv.Content.WriteTo(rc); err != nil {
			return err
		}
	}
}

// New creates a new instance of Proxy, which is either TLS encrypted or not.
func New(dst string, c net.Conn, s bool, ca *Cache) (*Proxy, error) {
	p := &Proxy{dst: dst, lc: c, cache: ca}
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
