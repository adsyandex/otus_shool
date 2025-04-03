package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization format"})
			return
		}

		claims, err := ParseToken(tokenString[7:])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}