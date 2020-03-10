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
	"strings"
	"sync"
	"time"

	"github.com/eclipse/paho.golang/packets"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
	"golang.org/x/sync/errgroup"
)

type Cache struct {
	Perms                  *sync.Map
	BanList                *sync.Map
	FailedTries            *sync.Map
	PubSummary, SubSummary map[string]map[string]bool
	SensitiveTopics        map[string]bool
}

type Proxy struct {
	dst    string
	lc, rc net.Conn
	pubKey common.Address
	*Cache
}

const (
	MAX_TRIES    = 2
	BAN_DURATION = time.Minute * 10
)

var (
	connackDenied = []byte{0x20, 0x03, 0x00, 0x87, 0x00}
	subackDenied  = []byte{0x90, 0x04, 0x00, 0x01, 0x00, 0x87}
	pubackDenied  = []byte{0x40, 0x04, 0x00, 0x01, 0x87, 0x00}
)

func (p *Proxy) savePub(topic, pubkey string) {
	if !p.Cache.SensitiveTopics[topic] {
		return
	}
	if _, ok := p.Cache.PubSummary[topic]; !ok {
		p.Cache.PubSummary[topic] = make(map[string]bool)
	}
	p.Cache.PubSummary[topic][pubkey] = true
}

func (p *Proxy) saveSub(topic, pubkey string) {
	if !p.Cache.SensitiveTopics[topic] {
		return
	}
	if _, ok := p.Cache.SubSummary[topic]; !ok {
		p.Cache.SubSummary[topic] = make(map[string]bool)
	}
	p.Cache.SubSummary[topic][pubkey] = true
}

func connIP(c net.Conn) string {
	requester := c.RemoteAddr().String()
	return requester[:strings.LastIndex(requester, ":")]
}

func (p *Proxy) applyFailedTry(topic string) {
	ip := connIP(p.lc)
	currentTries, _ := p.Cache.FailedTries.LoadOrStore(ip, 1)
	log.Printf("Auth failed for %q for topic %q", ip, topic)
	log.Printf("current tries %d", currentTries)
	if currentTries.(int) >= MAX_TRIES {
		log.Printf("Applying  ban for %q", ip)
		p.Cache.BanList.Store(ip, time.Now().Add(BAN_DURATION))
		p.Cache.FailedTries.Delete(ip)
		blockchain.PersistentLog(blockchain.Banned, p.pubKey, topic, connIP(p.lc))
		return
	}
	p.Cache.FailedTries.Store(ip, currentTries.(int)+1)
}

func (p *Proxy) verifyAccess(topic string, mask byte) (bool, byte, *sync.Map, error) {
	v, _ := p.Cache.Perms.LoadOrStore(topic, new(sync.Map))
	permMap, ok := v.(*sync.Map)
	if !ok {
		return false, 0, nil, fmt.Errorf("sync.Map not found in cache")
	}
	result, ok := permMap.LoadOrStore(p.pubKey, byte(0))
	resultByte, ok := result.(byte)
	if !ok {
		return false, 0, nil, fmt.Errorf("non-byte value found in permission map")
	}
	return ok && resultByte&mask == mask, resultByte, permMap, nil
}

func (p *Proxy) proxyPass(ctx context.Context, lc, rc net.Conn) error {
	defer ctx.Done()
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
				p.applyFailedTry("")
				return fmt.Errorf("flytrap signature not provided")
			}
			sig, err := base64.StdEncoding.DecodeString(sigEncoded)
			if err != nil {
				return err
			}
			pubCompressedEncoded, ok := conPack.Properties.User["flytrap_pub"]
			if !ok {
				p.applyFailedTry("")
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
				p.applyFailedTry("")
				return fmt.Errorf("public key signature check failed")
			}
			p.pubKey = crypto.PubkeyToAddress(*pub)
		// Check if it's SUBSCRIBE packet
		case recv.Type == packets.SUBSCRIBE:
			subPack := recv.Content.(*packets.Subscribe)
			for topic, _ := range subPack.Subscriptions {
				cached, resultByte, permMap, err := p.verifyAccess(topic, 2)
				if err != nil {
					return err
				}
				if cached {
					log.Printf("Client %q was already authorised to subscribe to topic %q", p.pubKey.String(), topic)
					p.saveSub(topic, p.pubKey.String())
					continue
				}
				ok, sensitive, err := blockchain.VerifyAccess(topic, connIP(p.lc), p.pubKey, country(connIP(p.lc)), false)
				if err != nil {
					return err
				}
				// Check with blockchain if pubkey is authorized to sub to given topic
				if ok {
					permMap.Store(p.pubKey, resultByte|2)
					log.Printf("Client %q was authorised to subscribe to topic %q", p.pubKey.String(), topic)
					p.Cache.SensitiveTopics[topic] = sensitive
					p.saveSub(topic, p.pubKey.String())
					continue
				}

				p.applyFailedTry(topic)
				log.Printf("Client %q attempted subscribe to topic %q and was denied due to insufficient permission", p.pubKey.String(), topic)
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
			cached, resultByte, permMap, err := p.verifyAccess(topic, 1)
			if err != nil {
				return err
			}
			if cached {
				log.Printf("Client %q was already authorised to publish to topic %q", p.pubKey.String(), topic)
				p.savePub(topic, p.pubKey.String())
				break
			}
			ok, sensitive, err := blockchain.VerifyAccess(topic, connIP(p.lc), p.pubKey, country(connIP(p.lc)), true)
			if err != nil {
				return err
			}
			// Check if pubkey is authorized to publish to given topic
			if !ok {
				log.Printf("Client %q attempted publish to topic %q and was denied due to insufficient permission", p.pubKey.String(), topic)
				p.applyFailedTry(topic)
				// Write PUBACK packet to client, denying access
				if _, err := lc.Write(pubackDenied); err != nil {
					return err
				}
				break
			}
			p.Cache.SensitiveTopics[topic] = sensitive
			permMap.Store(p.pubKey, resultByte|1)
			// Cache the result for future requests
			log.Printf("Client %q was authorised to publish to topic %q", p.pubKey.String(), topic)
			p.savePub(topic, p.pubKey.String())
		}
		if _, err := recv.Content.WriteTo(rc); err != nil {
			return err
		}
	}
}

// New creates a new instance of Proxy, which is either TLS encrypted or not.
func New(dst string, c net.Conn, s bool, ca *Cache) (*Proxy, error) {
	p := &Proxy{dst: dst, lc: c, Cache: ca}
	var err error

	// First check if user was banned before
	if val, ok := ca.BanList.Load(connIP(c)); ok {
		if !time.Now().After(val.(time.Time)) {
			c.Write(connackDenied)
			c.Close()
			return nil, fmt.Errorf("user was banned from connecting")
		}
	}

	// If requested TLS server, generate TLS connections
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
	// Otherwise use plain TCP server
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
