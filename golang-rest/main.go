package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "ok",
		})
	})
	server.Run(":4000")
}
