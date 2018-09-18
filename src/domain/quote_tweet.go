package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// QuoteTweet lslsls
type QuoteTweet struct {
	TextTweet
	quoteTweet *TextTweet
}

// NewQuoteTweet coment
func NewQuoteTweet(tweetUser, tweetText string, tweetQuote *TextTweet) *QuoteTweet {
	now := time.Now()
	id, _ := uuid.NewV4()

	return &QuoteTweet{TextTweet{id, tweetUser, tweetText, &now}, tweetQuote}
}

// PrintableTweet return a printable versión of tweet
func (tweet *QuoteTweet) PrintableTweet() string {
	return tweet.String()
}

func (tweet *QuoteTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text + " \"" + tweet.quoteTweet.PrintableTweet() + "\""
}
