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

const (
	TOKEN_LEN = 64
)

var (
	connIP    = flag.String("ip", "localhost:8888", "location of MQTT broker")
	authIP    = flag.String("auth", "localhost:8889", "location of auth server")
	connTls   = flag.Bool("tls", true, "whether to use TLS")
	pub       = flag.Bool("pub", false, "Whether to publish")
	sub       = flag.Bool("sub", false, "Whether to subscribe")
	pubMsg    = flag.String("msg", "Here Be Dragons", "message to be published")
	topic     = flag.String("topic", "MyTopic", "Topic for use for pub/sub")
	cID       = flag.String("id", "ClientID", "ID of connecting client")
	publicKey = flag.String("key", "0x641c46D43A3c552a76318c12D8Fc2839b913F32F", "Provide public key for Ethereum wallet")
)

func con(ip string) net.Conn {
	if !*connTls {
		con, err := net.Dial("tcp", ip)
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
	con, err := tls.Dial("tcp", ip, &tls.Config{RootCAs: c, InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}
	return con
}

func requestToken() (string, error) {
	c := con(*authIP)
	defer c.Close()
	c.Write([]byte("secret_password"))
	buf := make([]byte, TOKEN_LEN*2)
	n, err := c.Read(buf)
	if err != nil {
		return "", err
	}
	tok := string(buf[:n])
	if tok == "FAIL" {
		return "", fmt.Errorf("authentication failed")
	}
	return tok, nil
}

func subscribe(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	c := paho.NewClient(paho.ClientConfig{Conn: con(*connIP), Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		log.Printf("Message received: %s", string(m.Payload))
	})})
	// tok, err := requestToken()
	// if err != nil {
	// 	log.Printf("Token request failed with: %s. Continuing MQTT request without token.", err)
	// 	return
	// }
	_, err := c.Connect(ctx, &paho.Connect{
		ClientID:  *cID + "_SUB",
		KeepAlive: 5,
		Properties: &paho.ConnectProperties{
			User: map[string]string{
				"flytrap": *publicKey,
			}},
	})
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(&paho.Disconnect{})
	sub := paho.Subscribe{Subscriptions: map[string]paho.SubscribeOptions{
		*topic: paho.SubscribeOptions{QoS: 0},
	}}
	_, err = c.Subscribe(ctx, &sub)
	if err != nil {
		panic(err)
	}
	select {
	case <-ctx.Done():
		log.Print("Disconnecting Subscriber...")
		return
	}
}
func publish(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	c := paho.NewClient(paho.ClientConfig{Conn: con(*connIP)})
	tok, err := requestToken()
	if err != nil {
		log.Printf("Token request failed with: %s. Continuing MQTT request without token.", err)
		return
	}
	_, err = c.Connect(ctx, &paho.Connect{
		ClientID:  *cID + "_PUB",
		KeepAlive: 5,
		Properties: &paho.ConnectProperties{
			User: map[string]string{
				"flytrap": tok,
			}},
	})
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
		_, err = c.Publish(ctx, &paho.Publish{Topic: *topic, Payload: []byte(msg), QoS: 1})
		log.Printf("Published message: %s", msg)
		if err != nil {
			panic(err)
		}
		i += 1
		time.Sleep(time.Second)
	}
}
func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	if *pub {
		wg.Add(1)
		go publish(wg, ctx)
	}
	if *sub {
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