package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mysecret") // handlers と同じキーを使う

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Authorization ヘッダから取得
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(401, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		// Token をパース
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// ユーザー名(sub)を取り出せる
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user", claims["sub"])

		c.Next()
	}
}
