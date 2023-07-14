package repository

import (
	"fmt"
	"topicpad-api/db"
	"topicpad-api/model"
)

type UserRepository struct {}

type User model.User

type UserScan struct {
	ID uint `json:"id"`
	Auth0ID string `json:"auth0ID"`
	Kind int `json:"kind"`
	CreatedAt string `json:"createdAt"`
}

func (_ UserRepository) GetAll() ([]UserScan, error) {
	db := db.GetDB()
	var u []UserScan
	if err := db.Table("users").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (_ UserRepository) GetByID(auth0ID string) (UserScan, error) {
	db := db.GetDB()
	var u UserScan
	if err := db.Table("users").Where("auth0_id = ?", auth0ID).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (_ UserRepository) Create(auth0ID string, kind int) (UserScan, error) {
	db := db.GetDB()
	var u UserScan
	user := User{Auth0ID: auth0ID, Kind: kind}
	if err := db.Create(&user).Error; err != nil {
		return UserScan{}, err
	}
	fmt.Println("User作成成功!")
	if err := db.Table("users").Where("auth0_id = ?", auth0ID).First(&u).Error
		err != nil {
		return UserScan{}, err
	}
	return u, nil
}
