package repository

import (
	"fmt"
	"crypto/sha256"
	"topicpad-api/db"
	"topicpad-api/model"
)

type UserRepository struct {}

type User model.User

type UserScan struct {
	ID uint `json:"id"`
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

func (_ UserRepository) GetByID(ID uint) (UserScan, error) {
	db := db.GetDB()
	var u UserScan
	if err := db.Table("users").Where("id = ?", ID).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// func (_ UserRepository) Create(auth0ID string, kind int) (UserScan, error) {
// 	db := db.GetDB()
// 	user := User{Auth0ID: auth0ID, Kind: kind}
	// if err := db.Create(&user).Error; err != nil {
	// 	return UserScan{}, err
	// }
	// fmt.Println("User作成成功!")
// 	var u UserScan
// 	if err := db.Table("users").Where("auth0_id = ?", auth0ID).First(&u).Error
// 		err != nil {
// 		return UserScan{}, err
// 	}
// 	return u, nil
// }

func Encrypt(char string) string {
	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
	return encryptText
}

func (_ UserRepository) Create(name string, email string, password string) (UserScan, error) {
	db := db.GetDB()
	user := User{
			Name: name,
			Email: email,
			Password: Encrypt(password),
	}
	if err := db.Create(&user).Error; err != nil {
		return UserScan{}, err
	}
	fmt.Println("User作成成功!")
	var u UserScan
	if err := db.Table("users").Where("email = ?", email).First(&u).Error
		err != nil {
		return UserScan{}, err
	}

	return u, nil
}
