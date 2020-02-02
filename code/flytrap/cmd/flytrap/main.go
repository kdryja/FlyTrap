package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"

	"github.com/kdryja96/thesis/code/flytrap"
)

var (
	local_port  = flag.String("p", ":8888", "Port to use for proxy server")
	mqtt_broker = flag.String("b", "test.mosquitto.org:8883", "Provide address of MQTT Broker.")
	use_tls     = flag.Bool("tls", true, "Whether the proxy should be encrypted on both sides. You will need to provide .crt and .key files if so.")
	server_crt  = flag.String("crt", "server.crt", "Location of your server's .crt used for TLS connection")
	server_key  = flag.String("key", "server.key", "Location of your server's .key file used for TLS connection")
)

func server() (net.Listener, error) {
	if !*use_tls {
		l, err := net.Listen("tcp", *local_port)
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
	l, err := tls.Listen("tcp", *local_port, config)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func main() {
	flag.Parse()
	l, err := server()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Now accepting connections under %s", *local_port)
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		proxy, err := flytrap.New(*mqtt_broker, c, *use_tls)
		if err != nil {
			log.Print(err)
		}
		go proxy.Handle()
	}
}
