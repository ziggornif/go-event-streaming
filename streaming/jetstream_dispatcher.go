package streaming

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Event struct {
	MessageType string    `json:"messageType"`
	ID          string    `json:"id"`
	Message     string    `json:"message"`
	Date        time.Time `json:"date"`
	Author      string    `json:"author"`
}

type Dispatcher interface {
	Emit(subject string, event Event) error
}

type jetstreamDispatcher struct {
	js nats.JetStreamContext
}

const (
	streamName     = "TWITTERCLONE"
	streamSubjects = "TWITTERCLONE.*"
)

func NewJetStreamDispatcher() Dispatcher {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream() // Returns JetStreamContext
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{streamSubjects},
	})

	if err != nil {
		log.Println(err)
	}

	return &jetstreamDispatcher{js}
}

func (jsd *jetstreamDispatcher) Emit(subject string, event Event) error {
	jsonEvent, _ := json.Marshal(event)
	fmt.Println(fmt.Sprintf("%v.%v", streamName, subject))
	ack, err := jsd.js.Publish(fmt.Sprintf("%v.%v", streamName, subject), jsonEvent)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while publishing event (message-type: %v, id: %v)", event.MessageType, event.ID)
	}
	log.Printf("Event has been published %v\n", ack.Sequence)
	return nil
}
