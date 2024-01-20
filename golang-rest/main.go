package main

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"

	"github.com/omusegad/go-app/polymophism"
)




func main() {
	// server := gin.Default()
	// server.GET("/", func(c *gin.Context){
	// 	c.JSON(200, gin.H{
	// 		"message" : "ok",
	// 	})
	// })
	// server.Run(":4000")

	// initiate go person

	var c  polymophism.Shape = polymophism.Circle{}
	var s polymophism.Shape = polymophism.Square{}

	c.Rander()
	s.Rander()



}
