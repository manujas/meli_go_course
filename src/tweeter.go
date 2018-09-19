package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/gin-gonic/gin"
	"github.com/manujas/meli_go_course/src/domain"
	"github.com/manujas/meli_go_course/src/service"
)

var tweetManager *service.TweetManager

func main() {

	tweetManager = service.NewTweetManager()

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a tweet with an image",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the url of your image: ")

			url := c.ReadLine()

			tweet := domain.NewImageTweet(user, text, url)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuoteTweet",
		Help: "Publishes a tweet with a quote",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the id of the tweet you want to quote: ")

			id, _ := strconv.Atoi(c.ReadLine())

			quoteTweet := tweetManager.GetTweetById(id)

			tweet := domain.NewQuoteTweet(user, text, quoteTweet)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the id: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := tweetManager.GetTweetById(id)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			count := tweetManager.CountTweetsByUser(user)

			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			c.Println(tweets)

			return
		},
	})

	go startServer()
	shell.Run()

}

func startServer() {
	r := gin.Default()

	r.GET("/", getTweets)
	r.POST("/", createTweet)
	r.GET("/tweet/:id", getTweetById)
	r.GET("/tweets/:user", getTweetByUser)

	r.Run()
}

func getTweets(c *gin.Context) {
	c.JSON(http.StatusOK, tweetManager.GetTweets())
}

func createTweet(c *gin.Context) {
	var rawTweet domain.RawTweet
	if err := c.ShouldBindJSON(&rawTweet); err == nil {
		fmt.Println(rawTweet)
		tweet := createNewTweet(rawTweet)
		if _, err := tweetManager.PublishTweet(tweet); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "twitted",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error",
				"Msg":    "No se pudo crear el tweet",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"Msg":    err.Error(),
		})
	}

}

func createNewTweet(rt domain.RawTweet) domain.Tweet {
	switch rt.TweetType {
	case "image":
		return domain.NewImageTweet(rt.User, rt.Text, rt.Image)
	case "quote":
		quotedTweet := tweetManager.GetTweetById(rt.IdQuotedTweet)
		return domain.NewQuoteTweet(rt.User, rt.Text, quotedTweet)
	default:
		return domain.NewTextTweet(rt.User, rt.Text)
	}
}

func getTweetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"Msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, tweetManager.GetTweetById(id))
}

func getTweetByUser(c *gin.Context) {
	c.JSON(http.StatusOK, tweetManager.GetTweetsByUser(c.Param("user")))
}
