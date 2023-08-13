package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"topicpad-api/model/repository"
)

type AuthController struct{}

type InputUser struct {
	Name 	 string `form:"user[name]"`
	Email 	 string `form:"user[email]"`
	Password string `form:"user[password]"`
}

func (pc AuthController) Signup(c *gin.Context) {
	var r repository.UserRepository
	var user InputUser
	c.Bind(&user)
	p, err := r.Create(user.Name, user.Email, user.Password)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"user": p})
	}
}
