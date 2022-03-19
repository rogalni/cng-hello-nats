package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/rogalni/cng-hello-nats/config"
)

func main() {
	log.Println("Run consumer")
	nu := fmt.Sprintf("nats://%s", config.App.NatsUrl)
	opts := []nats.Option{nats.Name("cng-hello-nats-consumer")}
	nc, _ := nats.Connect(nu, opts...)

	var wg sync.WaitGroup
	wg.Add(50)
	nc.Subscribe("TEST", func(m *nats.Msg) {
		defer wg.Done()
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	wg.Wait()
}
