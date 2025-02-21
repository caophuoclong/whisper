package auth

import "github.com/gin-gonic/gin"

type Handlers interface {
	Health() gin.HandlerFunc
	Register() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	Login() gin.HandlerFunc
}
