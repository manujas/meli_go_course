package service

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/manujas/meli_go_course/src/domain"
)

// Tweet unico
var Tweet *domain.Tweet

// Tweets quiere comentario
var tweets []*domain.Tweet

//
var userTweetsMap map[string][]*domain.Tweet

// InitializeService prepare all to work
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	userTweetsMap = make(map[string][]*domain.Tweet)
}

// PublishTweet quiere un
func PublishTweet(tweet *domain.Tweet) (uuid.UUID, error) {
	if tweet.Text == "" {
		return tweet.ID, fmt.Errorf("text is required")
	}

	if tweet.User == "" {
		return tweet.ID, fmt.Errorf("user is required")
	}

	if len(tweet.Text) > 140 {
		return tweet.ID, fmt.Errorf("text exceeds 140 characters")
	}

	Tweet = tweet
	tweets = append(tweets, tweet)
	userTweetsMap[tweet.User] = append(userTweetsMap[tweet.User], tweet)
	return tweet.ID, nil
}

// GetTweet quiere un coment
func GetTweet() *domain.Tweet {
	return Tweet
}

// GetTweets quiere un coment
func GetTweets() []*domain.Tweet {
	return tweets
}

// GetTweetByID lalala
func GetTweetByID(id uuid.UUID) *domain.Tweet {
	for _, tweet := range tweets {
		if tweet.ID == id {
			return tweet
		}
	}

	return nil
}

// CountTweetsByUser get count
func CountTweetsByUser(user string) int {
	return len(GetTweetsByUser(user))
}

// GetTweetsByUser get count
func GetTweetsByUser(user string) []*domain.Tweet {
	return userTweetsMap[user]
}
