package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"candra/backend-api/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil JWT_SECRET setiap request, bukan dari variabel global
		jwtKey := []byte(config.GetEnv("JWT_SECRET", "secret_key"))
		authHeader := c.GetHeader("Authorization")
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token salah"})
			c.Abort()
			return
		}
		token := parts[1]

		claims := &jwt.RegisteredClaims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !tkn.Valid {
			fmt.Println("JWT error:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Set("username", claims.Subject)
		c.Next()
	}
}
