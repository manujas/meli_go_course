package service

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/manujas/meli_go_course/src/domain"
)

// TweetManager estructura general del manager
type TweetManager struct {
	tweet         domain.Tweet
	tweets        []domain.Tweet
	userTweetsMap map[string][]domain.Tweet
}

// NewTweetManager constructor
func NewTweetManager() *TweetManager {
	manager := new(TweetManager)
	manager.userTweetsMap = make(map[string][]domain.Tweet)
	return manager
}

// PublishTweet quiere un
func (manager *TweetManager) PublishTweet(tweet domain.Tweet) (uuid.UUID, error) {
	text := tweet.GetText()
	user := tweet.GetUser()

	if text == "" {
		return tweet.GetId(), fmt.Errorf("text is required")
	}

	if user == "" {
		return tweet.GetId(), fmt.Errorf("user is required")
	}

	if len(text) > 140 {
		return tweet.GetId(), fmt.Errorf("text exceeds 140 characters")
	}

	manager.tweet = tweet
	manager.tweets = append(manager.tweets, tweet)
	manager.userTweetsMap[user] = append(manager.userTweetsMap[user], tweet)
	return tweet.GetId(), nil
}

// GetTweet quiere un coment
func (manager *TweetManager) GetTweet() domain.Tweet {
	return manager.tweet
}

// GetTweets quiere un coment
func (manager TweetManager) GetTweets() []domain.Tweet {
	return manager.tweets
}

// GetTweetById lalala
func (manager *TweetManager) GetTweetById(id uuid.UUID) domain.Tweet {
	for _, tweet := range manager.tweets {
		if tweet.GetId() == id {
			return tweet
		}
	}

	return nil
}

// CountTweetsByUser get count
func (manager *TweetManager) CountTweetsByUser(user string) int {
	return len(manager.GetTweetsByUser(user))
}

// GetTweetsByUser get count
func (manager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return manager.userTweetsMap[user]
}
