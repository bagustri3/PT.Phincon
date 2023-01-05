package controllers

import (

	"github.com/bagustri3/PT.Phincon.git/initializer"
	"github.com/bagustri3/PT.Phincon.git/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var body struct {
		name string
	}
	c.Bind(&body)

	user := models.User{Name: body.name}

	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func PostAttend(c *gin.Context) {
	var body struct {
		UserId int
		Status bool
		Date   string
	}

	c.Bind(&body)

	attendance := models.Attendance{UserId: body.UserId, Date: body.Date, Status: body.Status}

	result := initializer.DB.Create(&attendance)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"attendance": attendance,
	})

}

func GetAttend(c *gin.Context) {

}