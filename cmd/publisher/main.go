package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rogalni/cng-hello-nats/config"
)

func main() {
	log.Println("Run publisher")
	nu := fmt.Sprintf("nats://%s", config.App.NatsUrl)
	opts := []nats.Option{nats.Name("cng-hello-nats-publisher")}
	nc, err := nats.Connect(nu, opts...)
	if err != nil {
		log.Fatalf("Error connect to nats: %v", err)
	}
	defer nc.Close()
	instance := rand.Intn(10)
	count := 100
	for i := 0; i < count; i++ {
		publish(nc, instance, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func publish(nc *nats.Conn, ic int, i int) {
	msg := fmt.Sprintf("Message from instance: %d, iteration: %d", ic, i)
	if err := nc.Publish("TEST", []byte(msg)); err != nil {
		log.Printf("Error publish: %v\n", err)
	}
}
