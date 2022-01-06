package streaming

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func NewJetStreamListener(evtChan chan Event) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("TWITTERCLONE.*", func(msg *nats.Msg) {
		msg.Ack()
		var event Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Fatal(err)
		}
		evtChan <- event

	}, nats.Durable("go-event-subscriber"), nats.ManualAck())
}
