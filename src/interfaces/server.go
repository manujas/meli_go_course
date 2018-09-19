package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manujas/meli_go_course/src/domain"
	"github.com/manujas/meli_go_course/src/service"
)

type ServerApp struct {
	tweetManager *service.TweetManager
}

func GetServerApp(tweetManager *service.TweetManager) *ServerApp {
	serverApp := new(ServerApp)
	serverApp.tweetManager = tweetManager

	return serverApp
}

func (serverApp *ServerApp) StartServer() {
	r := gin.Default()

	r.GET("/", serverApp.getTweets)
	r.POST("/", serverApp.createTweet)
	r.GET("/tweet/:id", serverApp.getTweetById)
	r.GET("/tweets/:user", serverApp.getTweetByUser)

	r.Run()
}

func (serverApp *ServerApp) getTweets(c *gin.Context) {
	c.JSON(http.StatusOK, serverApp.tweetManager.GetTweets())
}

func (serverApp *ServerApp) createTweet(c *gin.Context) {
	var rawTweet domain.RawTweet
	if err := c.ShouldBindJSON(&rawTweet); err == nil {
		fmt.Println(rawTweet)
		tweet := serverApp.createNewTweet(rawTweet)
		if _, err := serverApp.tweetManager.PublishTweet(tweet); err == nil {
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

func (serverApp *ServerApp) createNewTweet(rt domain.RawTweet) domain.Tweet {
	switch rt.TweetType {
	case "image":
		return domain.NewImageTweet(rt.User, rt.Text, rt.Image)
	case "quote":
		quotedTweet := serverApp.tweetManager.GetTweetById(rt.IdQuotedTweet)
		return domain.NewQuoteTweet(rt.User, rt.Text, quotedTweet)
	default:
		return domain.NewTextTweet(rt.User, rt.Text)
	}
}

func (serverApp *ServerApp) getTweetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"Msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, serverApp.tweetManager.GetTweetById(id))
}

func (serverApp *ServerApp) getTweetByUser(c *gin.Context) {
	c.JSON(http.StatusOK, serverApp.tweetManager.GetTweetsByUser(c.Param("user")))
}
