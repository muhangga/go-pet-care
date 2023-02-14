package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtToken = os.Getenv("JWT_TOKEN")

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id int64, email string) (string, error) {
	expiredTime := time.Now().Add(12 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"exp":     expiredTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtToken))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(jwtToken), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
