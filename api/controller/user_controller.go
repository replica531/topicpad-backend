package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"topicpad-api/model/repository"
)

type UserController struct{}

func (pc UserController) Index(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"users": p})
	}
}

func (pc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	p, err := u.GetByID(id)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"user": p})
	}
}
