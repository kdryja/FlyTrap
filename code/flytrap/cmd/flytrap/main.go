package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"sync"

	"github.com/kdryja/thesis/code/flytrap"
)

var (
	local_port  = flag.String("p", ":8888", "Port to use for proxy server")
	mqtt_broker = flag.String("b", "test.mosquitto.org:8883", "Provide address of MQTT Broker.")
	use_tls     = flag.Bool("tls", true, "Whether the proxy should be encrypted on both sides. You will need to provide .crt and .key files if so.")
	server_crt  = flag.String("crt", "server.crt", "Location of your server's .crt used for TLS connection")
	server_key  = flag.String("key", "server.key", "Location of your server's .key file used for TLS connection")
)

func server(a string) (net.Listener, error) {
	if !*use_tls {
		l, err := net.Listen("tcp", a)
		if err != nil {
			return nil, err
		}
		return l, nil
	}
	cer, err := tls.LoadX509KeyPair(*server_crt, *server_key)
	if err != nil {
		return nil, err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	l, err := tls.Listen("tcp", a, config)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func main() {
	flag.Parse()
	mainL, err := server(*local_port)
	if err != nil {
		log.Fatal(err)
	}
	defer mainL.Close()
	log.Printf("Now accepting connections under %s", *local_port)
	cache := flytrap.Cache{Perms: &sync.Map{}}
	for {
		c, err := mainL.Accept()
		if err != nil {
			log.Fatal(err)
		}
		proxy, err := flytrap.New(*mqtt_broker, c, *use_tls, &cache)
		if err != nil {
			log.Print(err)
		}
		go proxy.Handle()
	}
}
