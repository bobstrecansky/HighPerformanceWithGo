package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "bar",
		})
	})
	server.Run()
}
