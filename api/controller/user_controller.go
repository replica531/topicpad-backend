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
	auth0ID := c.Params.ByName("id")
	var u repository.UserRepository
	p, err := u.GetByID(auth0ID)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"user": p})
	}
}

type InputUser struct {
	Auth0ID string `form:"user[auth0ID]"`
	Kind    int    `form:"user[kind]"`
}

func (pc UserController) Create(c *gin.Context) {
	var u repository.UserRepository
	var user InputUser
	c.Bind(&user)
	p, err := u.Create(user.Auth0ID, user.Kind)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"user": p})
	}
}
