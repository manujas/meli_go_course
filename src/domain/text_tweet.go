package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TextTweet lslsls
type TextTweet struct {
	ID   uuid.UUID
	User string
	Text string
	Date *time.Time
}

// NewTextTweet coment
func NewTextTweet(tweetUser, tweetText string) *TextTweet {
	now := time.Now()
	id, _ := uuid.NewV4()

	return &TextTweet{id, tweetUser, tweetText, &now}
}

// PrintableTweet return a printable versión of tweet
func (tweet *TextTweet) PrintableTweet() string {
	return tweet.String()
}

func (tweet *TextTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text
}

// GetUser lala
func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

// GetText lala
func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

// GetDate lala
func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

// GetId lala
func (tweet *TextTweet) GetId() uuid.UUID {
	return tweet.ID
}
