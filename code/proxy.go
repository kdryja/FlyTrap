package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

const (
	LOCAL_PORT  = ":8888"
	MQTT_BROKER = "test.mosquitto.org:8883"
)

func proxyPass(ctx context.Context, lc, rc io.ReadWriteCloser) error {
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
		_, err = rc.Write(data)
		if err != nil {
			return err
		}
	}
}

func proxy(lc net.Conn) {
	defer lc.Close()
	fmt.Println("Starting new proxy connection")

	rootCA, err := ioutil.ReadFile("mosquitto.org.crt")
	if err != nil {
		fmt.Println(err)
		return
	}
	crt := x509.NewCertPool()
	crt.AppendCertsFromPEM(rootCA)

	rc, err := tls.Dial("tcp", MQTT_BROKER, &tls.Config{RootCAs: crt})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return proxyPass(ctx, lc, rc)
	})
	g.Go(func() error {
		return proxyPass(ctx, rc, lc)
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	l, err := tls.Listen("tcp", LOCAL_PORT, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go proxy(c)
	}
}
