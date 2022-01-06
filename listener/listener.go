package listener

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gitlab.com/ziggornif/go-event-streaming/streaming"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func makeListener(evtChan chan streaming.Event) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, w.Header())
		for {
			event := <-evtChan
			if err := conn.WriteJSON(event); err != nil {
				log.Println(err)
			}
		}
	}
}

func NewListener(router *gin.Engine) {
	var events []streaming.Event
	evtChan := make(chan streaming.Event)
	go streaming.NewJetStreamListener(evtChan)

	eventsListener := makeListener(evtChan)

	listenerRouter := router.Group("/listener")
	{
		listenerRouter.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, events)
			events = []streaming.Event{}
		})

		listenerRouter.GET("/ws", func(c *gin.Context) {
			eventsListener(c.Writer, c.Request)
		})
	}
}
