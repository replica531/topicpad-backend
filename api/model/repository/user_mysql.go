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
	CreatedAt string
}

func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Select("id", "kind", "created_at").Table("users").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (_ UserRepository) GetByID(id string) (UserProfile, error) {
	db := db.GetDB()
	var u UserProfile
	if err := db.Table("users").Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
