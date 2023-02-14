package middleware

import (
	"net/http"
	"strings"

	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
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

func (m *middleware) JWTMiddleware(userUsecase usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			formatter := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, formatter)
			return
		}

		if !strings.Contains(token, "Bearer") {
			formatter := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, formatter)
			return
		}

		tokenStr := ""
		arrayToken := strings.Split(token, " ")
		if len(arrayToken) == 2 {
			tokenStr = arrayToken[1]
		}

		tokenValid, err := ValidateToken(tokenStr)
		if err != nil {
			formatter := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, formatter)
			return
		}

		claim, ok := tokenValid.Claims.(jwt.MapClaims)

		if !ok || !tokenValid.Valid {
			formatter := utils.ErrorResponse("Unauthorized2", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, formatter)
			return
		}

		userID := int64(claim["user_id"].(float64))

		user, err := userUsecase.FindUserID(userID)
		if err != nil {
			formatter := utils.ErrorResponse("Unauthorized3", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, formatter)
			return
		}

		c.Set("currentUser", user)
	}
}
