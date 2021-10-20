package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/null-channel/news-feed/handlers"
)




var port string


func init (){
	flag.StringVar(&port,"port","8080","server port")
}


func main(){
	flag.Parse()
	r := gin.Default()
	r.GET("/api/news/feed",handlers.NewsFeed)
	r.GET("/api/news/brief",handlers.NewsBrief)
	r.GET("/api/legacy/feed",handlers.LegacyBreif)
	r.Run(":"+port)
}

