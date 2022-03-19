package main

import (
	"fmt"
	"log"
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
	time.Sleep(2 * time.Second)
	count := 50
	for i := 1; i < count+1; i++ {
		publish(nc, i)
		time.Sleep(1 * time.Second)
	}

}
func publish(nc *nats.Conn, i int) {
	msg := fmt.Sprintf("Welcome: %d", i)
	if err := nc.Publish("TEST", []byte(msg)); err != nil {
		log.Printf("Error publish: %v\n", err)
	}
}
