package handlers

import "github.com/gin-gonic/gin"

func HandlerHealth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

// GET /hello?name=Taro
func HandlerHello(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	c.JSON(200, gin.H{"message": "Hello " + name})
}

// POST /user でJSON受信
func HandlerPostUser(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}
	c.JSON(201, gin.H{"created": body})
}
