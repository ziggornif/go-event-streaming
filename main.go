package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/ziggornif/go-event-streaming/streaming"
	"gitlab.com/ziggornif/go-event-streaming/tweet"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	jsDispatcher := streaming.NewJetStreamDispatcher()

	router := gin.Default()
	router.Use(cors.Default())

	dsn := "host=localhost user=monitoring password=secret dbname=monitoring-article port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tweetService := tweet.NewTweetService(db, jsDispatcher)

	router.GET("/tweets", func(c *gin.Context) {
		tweets := tweetService.ListTweets()
		c.JSON(http.StatusOK, tweets)
	})

	router.POST("/tweets", func(c *gin.Context) {
		var input tweet.TweetRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tweets, createErr := tweetService.CreateTweet(input)
		if createErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
			return
		}

		c.JSON(http.StatusOK, tweets)
	})

	router.GET("/metrics", prometheusHandler())

	router.Run(":8080")
}
