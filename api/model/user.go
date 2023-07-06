package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Kind int `json:"kind"`
}
