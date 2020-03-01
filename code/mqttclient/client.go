package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
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
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	TOKEN_LEN = 64
)

var (
	connIP  = flag.String("ip", "localhost:8888", "location of MQTT broker")
	authIP  = flag.String("auth", "localhost:8889", "location of auth server")
	connTls = flag.Bool("tls", true, "whether to use TLS")
	pub     = flag.Int("pub", 0, "How many messages to publish")
	sub     = flag.Int("sub", 0, "How many messages to receive via subscription")
	pubMsg  = flag.String("msg", "Here Be Dragons", "message to be published")
	topic   = flag.String("topic", "MyTopic5", "Topic for use for pub/sub")
	cID     = flag.String("id", "ClientID", "ID of connecting client")
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

func signKey() (string, string, error) {
	priv, err := crypto.LoadECDSA("privkey.asc")
	if err != nil {
		return "", "", err
	}
	hash := crypto.PubkeyToAddress(priv.PublicKey).Hash().Bytes()
	sig, err := crypto.Sign(hash, priv)
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(sig), base64.StdEncoding.EncodeToString(crypto.CompressPubkey(&priv.PublicKey)), nil
}

func subscribe(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	pubCtx, cancel := context.WithCancel(context.Background())
	count := 0
	c := paho.NewClient(paho.ClientConfig{Conn: con(*connIP), Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		log.Printf("Message received: %s", string(m.Payload))
		if count += 1; count > *sub-1 {
			cancel()
		}
	})})
	sig, pubKey, err := signKey()
	if err != nil {
		log.Printf("Token request failed with: %s. Continuing MQTT request without token.", err)
	}
	_, err = c.Connect(ctx, &paho.Connect{
		ClientID:  *cID + "_SUB",
		KeepAlive: 10,
		Properties: &paho.ConnectProperties{
			User: map[string]string{
				"flytrap_sig": sig,
				"flytrap_pub": pubKey,
			}},
	})
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(&paho.Disconnect{})
	subPaho := paho.Subscribe{Subscriptions: map[string]paho.SubscribeOptions{
		*topic: paho.SubscribeOptions{QoS: 0},
	}}
	_, err = c.Subscribe(ctx, &subPaho)
	if err != nil {
		panic(err)
	}
	select {
	case <-ctx.Done():
		log.Print("Disconnecting Subscriber...")
	case <-pubCtx.Done():
		log.Printf("Received %d messages, as requested, disconnecting", *sub)
	}
}
func publish(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	c := paho.NewClient(paho.ClientConfig{Conn: con(*connIP)})
	sig, pubKey, err := signKey()
	if err != nil {
		log.Printf("Token request failed with: %s. Continuing MQTT request without token.", err)
	}
	_, err = c.Connect(ctx, &paho.Connect{
		ClientID:  *cID + "_PUB",
		KeepAlive: 10,
		Properties: &paho.ConnectProperties{
			User: map[string]string{
				"flytrap_sig": sig,
				"flytrap_pub": pubKey,
			}},
	})
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(&paho.Disconnect{})
	for i := 0; i < *pub; i++ {
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
		time.Sleep(time.Second)
	}
	log.Printf("Published %d messages, as requested, disconnecting", *pub)
}
func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	if *pub > 0 {
		wg.Add(1)
		go publish(wg, ctx)
	}
	if *sub > 0 {
		wg.Add(1)
		go subscribe(wg, ctx)
	}
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		select {
		case <-sigs:
			log.Print("Received SIGINT, disconnecting clients...")
			cancel()
		}
	}()
	wg.Wait()
}
