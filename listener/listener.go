package listener

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ziggornif/go-event-streaming/streaming"
	"net/http"
)

func NewListener(router *gin.Engine) {
	var events []string
	go streaming.NewJetStreamListener(&events)

	listenerRouter := router.Group("/listener")
	{
		listenerRouter.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, events)
			events = []string{}
		})
	}
}
