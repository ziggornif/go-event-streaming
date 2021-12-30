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
		router.LoadHTMLFiles("listener/listener.html")
		listenerRouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "listener.html", nil)
		})

		listenerRouter.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, events)
			events = []string{}
		})
	}
}
