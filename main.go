package main

import (
	"github.com/bagustri3/PT.Phincon.git/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message" : "asep",
		})
	})
	r.Run()
}