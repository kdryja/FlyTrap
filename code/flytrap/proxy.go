package flytrap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

type proxy struct {
	dst    string
	lc, rc net.Conn
}

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

func proxyPass(ctx context.Context, lc, rc net.Conn, p *proxy) error {
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
		log.Println(hex.Dump(data))
		// Check if it's CONNECT packet
		// if data[0] == 0x10 {
		// 	// Identify the position of two bytes indicating length of ClientID
		// 	// This is important, as the Protocal Name length varies - which is stored in 3rd and 4th byte.
		// 	i := 8 + (data[2]<<2 | data[3])
		// 	// Identify length of ClientID, which is the combination of two bytes
		// 	idLen := data[i]<<2 | data[i+1]
		// 	i += 2
		// 	// Strip username and password from the payload.
		// 	data = data[:i+idLen]
		// 	// Change the total-length byte with respect to the removed username/password.
		// 	data[1] = byte(len(data) - 2)
		// 	// Clear Password and Username Connect Flag.
		// 	data[i-5] = data[i-5] & 0x3F
		// }
		_, err = rc.Write(data)
		if err != nil {
			return err
		}
	}
}

// New creates a new instance of Proxy, which is either TLS encrypted or not.
func New(dst string, c net.Conn, s bool) (*proxy, error) {
	p := &proxy{dst: dst, lc: c}
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
func (p *proxy) Handle() {
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
