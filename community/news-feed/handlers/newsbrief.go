package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/null-channel/news-feed/models"
)




func NewsBrief(c *gin.Context){
	var brief models.NewsArticle
	c.JSON(200,brief.Breif())
} 
