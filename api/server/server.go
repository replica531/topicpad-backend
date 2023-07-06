package server

import (
	"github.com/gin-gonic/gin"
	"topicpad-api/controller"
)

func Init() {
	router := router()
	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	users := router.Group("/users")
	{
		ctrl := controller.UserController{}
		users.GET("", ctrl.Index)
		// users.GET("/:id", ctrl.UserShow)
		// users.POST("/", ctrl.UserCreate)
		// users.PUT("/:id", ctrl.UserUpdate)
		// users.DELETE("/:id", ctrl.UserDelete)
	}

	return router
}
