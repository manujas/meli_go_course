package domain

import "time"

// Tweet structura
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

// NewTweet coment
func NewTweet(tweetUser, tweetText string) *Tweet {
	now := time.Now()
	return &Tweet{tweetUser, tweetText, &now}
}
