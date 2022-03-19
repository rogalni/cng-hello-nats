package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/rogalni/cng-hello-nats/config"
)

func main() {
	log.Println("Started subscriber")
	nu := fmt.Sprintf("nats://%s", config.App.NatsUrl)
	opts := []nats.Option{nats.Name("cng-hello-nats-consumer")}
	nc, _ := nats.Connect(nu, opts...)
	defer nc.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	nc.Subscribe("TEST", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	wg.Wait()
}
