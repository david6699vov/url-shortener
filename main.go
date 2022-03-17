package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-url-shortener/handler"
	"go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hey go URL Shortener !",
		})
	})

	//create short url
	r.POST("/create-short-url", func(context *gin.Context) {
		handler.CreateShortUrl(context)
	})

	r.GET("/:shortUrl", func(context *gin.Context) {
		handler.HandleShortUrlRedirect(context)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
