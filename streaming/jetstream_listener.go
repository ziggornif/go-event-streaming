package streaming

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

func NewJetStreamListener(events *[]Event) {
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
		*events = append(*events, event)

	}, nats.Durable("go-event-subscriber"), nats.ManualAck())
}
