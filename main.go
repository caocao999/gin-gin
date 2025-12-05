package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	myroutes(r)

	// // Health check
	// r.GET("/health", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"status": "ok"})
	// })

	// // GET /hello?name=Taro
	// r.GET("/hello", func(c *gin.Context) {
	// 	name := c.Query("name")
	// 	if name == "" {
	// 		name = "World"
	// 	}
	// 	c.JSON(200, gin.H{"message": "Hello " + name})
	// })

	// // POST /user でJSON受信
	// r.POST("/user", func(c *gin.Context) {
	// 	var body struct {
	// 		Name string `json:"name"`
	// 		Age  int    `json:"age"`
	// 	}
	// 	if err := c.BindJSON(&body); err != nil {
	// 		c.JSON(400, gin.H{"error": "invalid json"})
	// 		return
	// 	}
	// 	c.JSON(201, gin.H{"created": body})
	// })

	r.Run(":8080")
}
