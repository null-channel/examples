package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/null-channel/news-feed/models"
)


func NewsFeed( c *gin.Context  ) {
	var articles models.NewsArticle
	c.JSON(200,articles.GenArticles(10))
}

func LegacyBreif( c *gin.Context) {
	var articles models.NewsArticle
	c.XML(200,articles.GenArticles(5))
}


