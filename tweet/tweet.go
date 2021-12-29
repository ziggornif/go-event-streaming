package tweet

import (
	"gitlab.com/ziggornif/go-event-streaming/streaming"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Message string
	Author  string
}

func (t *Tweet) ToResponse() TweetResponse {
	return TweetResponse{
		ID:        t.ID,
		Message:   t.Message,
		CreatedAt: t.CreatedAt,
		Author:    t.Author,
	}
}

type TweetRequest struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

func (t *TweetRequest) NewTweet() Tweet {
	id, _ := uuid.NewUUID()
	return Tweet{
		ID:      id.String(),
		Message: t.Message,
		Author: t.Author,
	}
}

type TweetResponse struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
}

type tweetService struct {
	dbConn              *gorm.DB
	streamingDispatcher streaming.Dispatcher
}

type TweetService interface {
	ListTweets() []TweetResponse
	CreateTweet(request TweetRequest) (*TweetResponse, error)
}

func NewTweetService(dbConn *gorm.DB, streaming streaming.Dispatcher) TweetService {
	dbConn.AutoMigrate(&Tweet{})

	return &tweetService{
		dbConn,
		streaming,
	}
}

func (ts *tweetService) ListTweets() []TweetResponse {
	var tweets []Tweet
	ts.dbConn.Order("created_at DESC").Limit(50).Find(&tweets)

	var results []TweetResponse
	for _, tweet := range tweets {
		results = append(results, tweet.ToResponse())
	}
	return results
}

func (ts *tweetService) CreateTweet(request TweetRequest) (*TweetResponse, error) {
	tweet := request.NewTweet()
	ts.dbConn.Create(&tweet)
	err := ts.streamingDispatcher.Emit("tweet_created", streaming.Event{
		MessageType: "tweet_created",
		ID:          tweet.ID,
		Message:     tweet.Message,
		Date:        tweet.CreatedAt,
		Author:      tweet.Author,
	})
	if err != nil {
		return nil, err
	}

	tweetResp := tweet.ToResponse()
	return &tweetResp, nil
}
