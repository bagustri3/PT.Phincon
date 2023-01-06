package main

import (
	"github.com/bagustri3/PT.Phincon.git/initializer"
	"github.com/bagustri3/PT.Phincon.git/controller"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	
	r.POST("/user", controllers.AddUser)

	r.POST("/attend", controllers.PostAttend)
	r.GET("/attend/:id", controllers.GetAttend)

	r.Run()
}