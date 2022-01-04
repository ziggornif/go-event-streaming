package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/ziggornif/go-event-streaming/listener"
	"gitlab.com/ziggornif/go-event-streaming/streaming"
	"gitlab.com/ziggornif/go-event-streaming/tweet"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	jsDispatcher := streaming.NewJetStreamDispatcher()

	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/public", "./public")

	dsn := "host=localhost user=monitoring password=secret dbname=monitoring-article port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Init events listener
	listener.NewListener(router)

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

	router.POST("/tweets/:id/likes", func(c *gin.Context) {
		tweetID := c.Param("id")
		likeErr := tweetService.LikeTweet(tweetID)
		if likeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": likeErr.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	})

	router.LoadHTMLFiles("./public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":8080")
}
