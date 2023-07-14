package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"

	"topicpad-api/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
        "http://localhost:3000",
    },
    // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
    AllowMethods: []string{
				"PUT",
        "POST",
        "GET",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
				"X-CSRF-Token",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  }))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "secret531",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	router.GET("/csrf", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token": csrf.GetToken(c),
		})
	})

	users := router.Group("/users")
	{
		ctrl := controller.UserController{}
		users.GET("", ctrl.Index)
		users.GET("/:id", ctrl.Show)
		users.POST("", ctrl.Create)
		// users.PUT("/:id", ctrl.UserUpdate)
		// users.DELETE("/:id", ctrl.UserDelete)
	}

	return router
}
