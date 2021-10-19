package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/null-channel/user-service/models"
)


func GetUsers( c *gin.Context ){
       var user models.User 
       c.JSON(200,user.GenUsers(10))
}
