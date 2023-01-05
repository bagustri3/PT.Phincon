package main

import (
	"github.com/bagustri3/PT.Phincon.git/initializer"
	"github.com/bagustri3/PT.Phincon.git/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.User{})
	initializer.DB.AutoMigrate(&models.Attendance{})

}
