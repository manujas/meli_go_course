package domain

type RawTweet struct {
	User          string `json:"user"`
	Text          string `json:"text"`
	Image         string `json:"image"`
	IdQuotedTweet int    `json:"idQuotedTweet"`
	TweetType     string `json:"tweetType"`
}
