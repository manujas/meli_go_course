package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ImageTweet lslsls
type ImageTweet struct {
	TextTweet
	image string
}

// NewImageTweet coment
func NewImageTweet(tweetUser, tweetText, tweetImage string) *ImageTweet {
	now := time.Now()
	id, _ := uuid.NewV4()

	return &ImageTweet{TextTweet{id, tweetUser, tweetText, &now}, tweetImage}
}

// PrintableTweet return a printable versión of tweet
func (tweet *ImageTweet) PrintableTweet() string {
	return tweet.String()
}

func (tweet *ImageTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text + " " + tweet.image
}
