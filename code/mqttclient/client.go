package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	PROXY = "ssl://localhost:8888"
)

func options(id string) *mqtt.ClientOptions {
	pem, _ := ioutil.ReadFile("server.crt")
	crt := x509.NewCertPool()
	crt.AppendCertsFromPEM(pem)
	tlsOpts := tls.Config{RootCAs: crt, InsecureSkipVerify: true}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(PROXY)
	opts.SetClientID(id)
	opts.SetTLSConfig(&tlsOpts)
	return opts
}

func subscribe() {
	choke := make(chan [2]string)
	opts := options("subscriber")
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})
	opts.SetClientID("gotrivial")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("listening for messages")
	incoming := <-choke
	fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	c.Disconnect(250)
}

func main() {
	go subscribe()
	opts := options("publisher")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	time.Sleep(1 * time.Second)

	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)
	c.Disconnect(250)
}
