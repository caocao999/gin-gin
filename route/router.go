package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/gin-gin/handlers"
)

func MyRoutes(r *gin.Engine) {

	// Health check
	r.GET("/health", handlers.HandlerHealth)

	// GET /hello?name=Taro
	r.GET("/hello", handlers.HandlerHello)

	// POST /user でJSON受信
	r.POST("/user", handlers.HandlerPostUser)
}
