package controller

import (
	"net/http"
	"strconv"

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
	userId, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	var u repository.UserRepository
	p, err := u.GetByID(uint(userId))
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"user": p})
	}
}
