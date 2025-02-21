package utils

import (
	"time"

	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func GenerateJWTToken(user *models.User, config *configs.Config, tokenType string) (string, error) {
	var expired int64
	switch tokenType {
	case "access":
		expired = time.Now().Add(time.Minute * 60).Unix()
	case "refresh":
		expired = time.Now().Add(time.Minute * 60 * 24 * 30).Unix()
	default:
		expired = time.Now().Add(time.Minute * 60).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   expired,
	})
	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
