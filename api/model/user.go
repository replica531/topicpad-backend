package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"uniqueIndex:idx_user_name;size:100;unique;not null"`
	Email string `json:"email" binding:"required" gorm:"uniqueIndex:idx_user_email;size:100;unique;not null"`
	Password string `json:"password" binding:"required" gorm:"size:100;not null"`
	Kind int `json:"kind" binding:"required" gorm:"default:0"`
	Topics []Topic
}

// 0: admin
// 1: member
