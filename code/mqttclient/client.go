package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

var (
	connIP  = flag.String("ip", "localhost:8888", "location of MQTT broker")
	connTls = flag.Bool("tls", true, "whether to use TLS")
	pubMsg  = flag.String("msg", "", "message to be published")
	pubTpc  = flag.String("topic", "MyTopic", "topic to publish messages to")
	subTpc  = flag.String("sub", "", "topic to subscribe to")
	cID     = flag.String("id", "ClientID", "ID of connecting client")
)

func con() net.Conn {
	if !*connTls {
		con, err := net.Dial("tcp", *connIP)
		if err != nil {
			panic(err)
		}
		return con
	}
	c := x509.NewCertPool()
	pem, err := ioutil.ReadFile("server.crt")
	if err != nil {
		panic(err)
	}
	c.AppendCertsFromPEM(pem)
	con, err := tls.Dial("tcp", *connIP, &tls.Config{RootCAs: c, InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}
	return con
}

func subscribe(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	c := paho.NewClient(paho.ClientConfig{Conn: con(), Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		log.Printf("Message received: %s", string(m.Payload))
	})})
	_, err := c.Connect(ctx, &paho.Connect{ClientID: *cID + "_SUB", KeepAlive: 10})
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(&paho.Disconnect{})
	sub := paho.Subscribe{Subscriptions: map[string]paho.SubscribeOptions{
		*subTpc: paho.SubscribeOptions{QoS: 0},
	}}
	_, err = c.Subscribe(ctx, &sub)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 3)
	return
	select {
	case <-ctx.Done():
		log.Print("Disconnecting Subscriber...")
		return
	}
}
func publish(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	c := paho.NewClient(paho.ClientConfig{Conn: con()})
	_, err := c.Connect(ctx, &paho.Connect{ClientID: *cID + "_PUB", KeepAlive: 5})
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(&paho.Disconnect{})
	i := 1
	for {
		select {
		case <-ctx.Done():
			log.Print("Disconnecting Publisher...")
			return
		default:
		}
		msg := fmt.Sprintf("%s #%d", *pubMsg, i)
		_, err = c.Publish(ctx, &paho.Publish{Topic: *pubTpc, Payload: []byte(msg), QoS: 1})
		log.Printf("Published message: %s", msg)
		if err != nil {
			panic(err)
		}
		i += 1
		time.Sleep(time.Second)
		return
	}
}
func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	if *pubMsg != "" {
		wg.Add(1)
		go publish(wg, ctx)
	}
	if *subTpc != "" {
		wg.Add(1)
		go subscribe(wg, ctx)
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	<-sigs
	log.Print("Received SIGINT, disconnecting clients...")
	cancel()
	wg.Wait()

}
