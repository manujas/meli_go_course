package service

import (
	"fmt"

	"github.com/manujas/meli_go_course/src/domain"
)

// Tweet unico
var Tweet *domain.Tweet

// Tweets quiere comentario
var tweets []*domain.Tweet

// InitializeService prepare all to work
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
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

	Tweet = tweet
	tweets = append(tweets, tweet)
	return nil
}

// GetTweet quiere un coment
func GetTweet() *domain.Tweet {
	return Tweet
}

// GetTweets quiere un coment
func GetTweets() []*domain.Tweet {
	return tweets
}
