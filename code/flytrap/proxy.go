package flytrap

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

type Proxy struct {
	dst    string
	lc, rc net.Conn
}

const PROP = "flytrap"

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
			log.Print(hex.Dump(data))
			sep := []byte(PROP)
			// Located flytrap specific property in the payload.
			i := bytes.Index(data, sep)
			if i == -1 {
				log.Print("token not provided!")
			} else {
				// Location of the secret key / value property.
				loc := i + len(sep)
				keyLen := binary.BigEndian.Uint16(data[loc : loc+2])
				token := string(data[loc+2 : loc+2+int(keyLen)])
				log.Printf("Token has been found: %s", token)
			}
		}
		_, err = rc.Write(data)
		if err != nil {
			return err
		}
	}
}

// New creates a new instance of Proxy, which is either TLS encrypted or not.
func New(dst string, c net.Conn, s bool) (*Proxy, error) {
	p := &Proxy{dst: dst, lc: c}
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
