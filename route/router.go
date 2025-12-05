package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/gin-gin/handlers"
	"github.com/yourname/gin-gin/middleware"
)

func MyRoutes(r *gin.Engine) {

	// Health check
	r.GET("/health", handlers.HandlerHealth)

	// GET /hello?name=Taro
	r.GET("/hello", handlers.HandlerHello)

	// POST /user でJSON受信
	r.POST("/user", handlers.HandlerPostUser)

	r.POST("/login", handlers.HandlerLogin)

	// JWT 必須のグループ
	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/secure", handlers.HandlerSecure)
	}
}
