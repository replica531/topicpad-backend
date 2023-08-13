package repository

import (
	"fmt"
	"topicpad-api/db"
	"topicpad-api/model"
)

type TopicRepository struct {}

type Topic model.Topic

type TopicScan struct {
	ID uint `json:"id"`
	Content string `json:"content"`
	Private bool `json:"private"`
	UserID uint `json:"user_id"`
	CreatedAt string `json:"createdAt"`
}

func (_ TopicRepository) GetAll() ([]TopicScan, error) {
	db := db.GetDB()
	var t []TopicScan
	if err := db.Table("topics").Scan(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (_ TopicRepository) GetByID(id uint) (TopicScan, error) {
	db := db.GetDB()
	var t TopicScan
	if err := db.Table("topics").Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (_ TopicRepository) Create(content string, private bool, userID uint) (TopicScan, error) {
	db := db.GetDB()
	topic := Topic{Content: content, Private: private, UserID: userID}
	if err := db.Create(&topic).Error; err != nil {
		return TopicScan{}, err
	}
	fmt.Println("Topic作成成功!")
	var t TopicScan
	if err := db.Table("topics").Where("id = ?", topic.ID).First(&t).Error; err != nil {
		return TopicScan{}, err
	}
	return t, nil
}
