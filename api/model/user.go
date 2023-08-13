package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Auth0ID string `json:"auth0_id" binding:"required" gorm:"uniqueIndex:idx_auth0_id;size:100"`
	Kind int `json:"kind" binding:"required" gorm:"default:0"`
	Topics []Topic
}

// 0: admin
// 1: member
