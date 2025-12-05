package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/gin-gin/route"
)

func main() {
	r := gin.Default()

	route.MyRoutes(r)

	r.Run(":8080")
}
