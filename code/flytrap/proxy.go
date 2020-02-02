package flytrap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
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

func proxyPass(ctx context.Context, lc, rc net.Conn, p *proxy) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		b := make([]byte, 4096)
		n, err := lc.Read(b)
		if err != nil {
			return err
		}
		data := make([]byte, n)
		copy(data, b[:n])
		if p.lc.RemoteAddr() == lc.RemoteAddr() {
			//TODO(kdryja): Process data - which is TCP packet
			log.Print("Left side of the proxy")
		}
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
