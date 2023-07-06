package repository

import (
	"topicpad-api/db"
	"topicpad-api/model"
)

type UserRepository struct {}

type User model.User

type UserProfile struct {
	ID uint
	Kind int
}

func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Select("id", "kind").Table("users").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

