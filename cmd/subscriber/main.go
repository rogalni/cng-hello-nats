package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/rogalni/cng-hello-nats/api/model"
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

	subscribeNats(nc)
	subscribeJetStream(nc)

	// Wait forever
	wg.Wait()
}
func subscribeJetStream(nc *nats.Conn) {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("MESSAGE.*", func(msg *nats.Msg) {
		msg.Ack()
		var m model.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("cng-hello-nats-subscriber service subscribes from subject:%s\n", msg.Subject)
		log.Printf("MessageId: %d, MessageText: %s\n", m.Id, m.Text)
	}, nats.Durable("cng-hello-nats-subscriber"), nats.ManualAck())

}

func subscribeNats(nc *nats.Conn) {
	nc.Subscribe("TEST", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
}
