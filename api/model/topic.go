package model

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Content string `json:"content" binding:"required" gorm:"uniqueIndex:idx_topic_content;size:100"`
	Private bool `json:"private" binding:"required" gorm:"default:false"`
	UserID uint `json:"user_id" binding:"required"`
	Tags []Tag `gorm:"many2many:topic_tags;"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"uniqueIndex:idx_tag_name;size:100"`
	Topics []Topic `gorm:"many2many:topic_tags;"`
}
