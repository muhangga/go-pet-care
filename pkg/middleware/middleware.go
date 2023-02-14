package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type middleware struct{}

func InitMiddleware() *middleware {
	return &middleware{}
}

func (m *middleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, client-security-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Request-Headers", "Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (m *middleware) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if !strings.Contains(token, "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		tokenStr := ""
		arrayToken := strings.Split(token, " ")
		if len(arrayToken) == 2 {
			tokenStr = arrayToken[1]
		}

		claims, err := ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("user_id", claims)
		c.Next()
	}
}
