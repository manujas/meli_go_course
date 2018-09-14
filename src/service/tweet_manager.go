package service

// Tweet quiere comentario
var Tweet string

// PublishTweet quiere un
func PublishTweet(tweet string) {
	Tweet = tweet
}

// GetTweet quiere un coment
func GetTweet() string {
	return Tweet
}
