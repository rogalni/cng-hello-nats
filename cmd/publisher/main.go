package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rogalni/cng-hello-nats/api/model"
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
	// instance := rand.Intn(10)

	// runNats(nc, instance)
	runJetStream(nc)

}

func runJetStream(nc *nats.Conn) {
	js, err := nc.JetStream()
	if err != nil {
		fmt.Printf("Error init JetStream")
	}
	createStream(js)

	for i := 0; i < 100; i++ {

		msg := &model.Message{
			Id:   i,
			Text: "Example mesage",
		}
		msgj, _ := json.Marshal(msg)

		_, err = js.Publish("MESSAGE.TEST", msgj)
		if err != nil {
			fmt.Printf("Error publish to stream: %v", err)
		}
	}
}

func runNats(nc *nats.Conn, instance int) {
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

const (
	streamName     = "MESSAGE"
	streamSubjects = "MESSAGE.*"
)

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the MESSAGE stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
