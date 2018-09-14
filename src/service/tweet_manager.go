package service

import "github.com/manujas/meli_go_course/src/domain"

// Tweet quiere comentario
var Tweet *domain.Tweet

// PublishTweet quiere un
func PublishTweet(tweet *domain.Tweet) {
	Tweet = tweet
}

// GetTweet quiere un coment
func GetTweet() *domain.Tweet {
	return Tweet
}
