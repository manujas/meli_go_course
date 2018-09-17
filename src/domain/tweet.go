package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Tweet structura
type Tweet struct {
	ID   uuid.UUID
	User string
	Text string
	Date *time.Time
}

// NewTweet coment
func NewTweet(tweetUser, tweetText string) *Tweet {
	now := time.Now()
	id, _ := uuid.NewV4()

	return &Tweet{id, tweetUser, tweetText, &now}
}

// PrintableTweet return a printable versión of tweet
func (tweet *Tweet) PrintableTweet() string {
	return tweet.String()
}

func (tweet *Tweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text
}
