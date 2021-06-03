package controllers

import (
	"gin-first/platform/newsFeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed newsFeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
