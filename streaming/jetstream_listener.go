package streaming

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func NewJetStreamListener(events *[]string) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("TWITTERCLONE.tweet_created", func(msg *nats.Msg) {
		msg.Ack()
		var event Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Fatal(err)
		}
		*events = append(*events, fmt.Sprintf("%v (author %v)", event.Message, event.Author))

	}, nats.Durable("go-event-subscriber"), nats.ManualAck())
}
