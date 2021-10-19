package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/null-channel/user-service/models"
)



func GetUser( c *gin.Context ){
	u := models.User{}
  	c.JSON(200,u.NewUser()) 
}
