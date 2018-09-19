package main

import (
	"github.com/manujas/meli_go_course/src/interfaces"
	"github.com/manujas/meli_go_course/src/service"
)

var tweetManager *service.TweetManager

func main() {
	tweetManager = service.NewTweetManager()

	shellApp := interfaces.GetShellApp(tweetManager)
	serverApp := interfaces.GetServerApp(tweetManager)

	go serverApp.StartServer()
	shellApp.Shell.Run()
}
