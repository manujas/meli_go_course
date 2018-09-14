package service

import (
	"fmt"

	"github.com/manujas/meli_go_course/src/domain"
)

// Tweet quiere comentario
var Tweet *domain.Tweet

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

	Tweet = tweet
	return nil
}

// GetTweet quiere un coment
func GetTweet() *domain.Tweet {
	return Tweet
}
