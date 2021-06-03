package controllers

import (
	"gin-first/platform/newsFeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsFeedRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsFeedPost(feed newsFeed.Setter) gin.HandlerFunc {
	// for testing in browser
	// await fetch("/newsfeed", {
	// 	method: "POST",
	// 	headers: { "Content-type": "application/json" },
	// 	body: JSON.stringify({
	// 		title: "you",
	// 		post: "description"
	// 	})

	// })

	return func(c *gin.Context) {
		requestBody := newsFeedRequest{}
		c.Bind(&requestBody)

		item := newsFeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
