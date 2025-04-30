package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"uniqueIndex;not null"`
	Email      string `json:"email" gorm:"uniqueIndex;not null"`
	Password   string `json:"password" gorm:"not null"`
	SlackID    string `json:"slack_id" gorm:"default:''"`
	Experience int    `json:"experience" gorm:"default:0"`
}
