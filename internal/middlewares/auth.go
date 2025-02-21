package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (mw *MiddlewareManager) AuthJWTFromRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerString := c.GetHeader("Authorization")
		if bearerString == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "No token provided",
				},
			)
			return
		}
		tokenParts := strings.Split(bearerString, " ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Token is invalid",
				},
			)
			return
		}
		token := tokenParts[1]
		if err := mw.validateJWTToken(token, c); err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				err.Error(),
			)
		}
		c.Next()
	}
}

func (mw *MiddlewareManager) validateJWTToken(tk string, c *gin.Context) error {
	token, err := jwt.Parse(
		tk,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mw.cfg.JWT.Secret), nil
		})
	if err != nil {
		return err
	}
	if !token.Valid {
		mw.logger.Info("Token is invalid")
		return errors.New("token is invalid")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, ok := claims["id"].(string)
		if !ok {
			return errors.New("token is invalid")
		}
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			return err
		}
		user, err := mw.au.GetUserById(c, userUUID)
		if err != nil {
			return err
		}
		user.EmptyPassword()
		c.Set("user", user)
	}
	return nil
}
