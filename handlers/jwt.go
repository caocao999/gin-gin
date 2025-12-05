package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mysecret") // 本来は環境変数

// ログイン：正しいユーザー名/パスワードならJWTを返す
func HandlerLogin(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	// 最小版：実際のDB認証ではない
	if body.Username != "admin" || body.Password != "pass" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// JWT ペイロード
	claims := jwt.MapClaims{
		"sub": body.Username,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	// Token 作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "token error"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func HandlerSecure(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Authorized!",
	})
}
