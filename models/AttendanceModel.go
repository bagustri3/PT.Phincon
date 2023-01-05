package models

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	UserId int    `json:"user_id"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}