package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"golang.org/x/sync/errgroup"
)

func proxyPass(ctx context.Context, lc, rc io.ReadWriteCloser) error {
	for {
		b := make([]byte, 4096)
		n, err := lc.Read(b)
		if err != nil {
			return err
		}
		data := make([]byte, n)
		copy(data, b)
		_, err = rc.Write(data)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Millisecond * 1000):
		}
	}
}

func proxy(lc net.Conn) {
	rc, err := net.Dial("tcp", ":1883")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()
	defer lc.Close()
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return proxyPass(ctx, lc, rc)
	})
	g.Go(func() error {
		return proxyPass(ctx, rc, lc)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Proxy finished!\n")
		fmt.Println(err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go proxy(c)
	}
}
