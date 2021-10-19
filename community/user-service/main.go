package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/null-channel/user-service/handlers"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "server port")
}


func main(){
	flag.Parse()
	r := gin.Default()
	r.GET("/api/user",handlers.GetUser)
	r.GET("/api/users",handlers.GetUsers)
	r.Run(":"+port)
}
