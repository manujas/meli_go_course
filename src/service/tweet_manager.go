package service

import (
	"fmt"

	"github.com/manujas/meli_go_course/src/domain"
)

// Tweets quiere comentario
var Tweets []*domain.Tweet

// InitializeService prepare all to work
func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
}

// PublishTweet quiere un
func PublishTweet(tweet *domain.Tweet) error {
	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}

	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}

	if len(tweet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	Tweets = append(Tweets, tweet)
	return nil
}

// GetTweets quiere un coment
func GetTweets() []*domain.Tweet {
	return Tweets
}
