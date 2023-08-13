package controller

import (
	"net/http"
	"strconv"

	"topicpad-api/model/repository"

	"github.com/gin-gonic/gin"
)

type TopicController struct{}

func (pc TopicController) Index(c *gin.Context) {
	var t repository.TopicRepository
	p, err := t.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"topics": p})
	}
}

func (pc TopicController) Show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	var t repository.TopicRepository
	p, err := t.GetByID(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"topic": p})
	}
}

type InputTopic struct {
	Content string `form:"topic[content]"`
	Private bool   `form:"topic[private]"`
	UserID  uint   `form:"topic[userID]"`
}

func (pc TopicController) Create(c *gin.Context) {
	var t repository.TopicRepository
	var topic InputTopic
	c.Bind(&topic)
	p, err := t.Create(topic.Content, topic.Private, topic.UserID)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"topic": p})
	}
}
