package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("start server...")
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	log.Fatal(r.Run()) // default :8080
}
