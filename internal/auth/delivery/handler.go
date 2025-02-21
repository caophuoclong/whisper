package delivery

import (
	"fmt"
	"net/http"

	"github.com/caophuoclong/whisper/internal/auth"
	"github.com/caophuoclong/whisper/internal/models"
	"github.com/caophuoclong/whisper/pkg"
	"github.com/caophuoclong/whisper/pkg/httpErrors"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	uc     auth.AuthUsecase
	logger pkg.Logger
}

func NewAuthHandler(uc auth.AuthUsecase, logger pkg.Logger) auth.Handlers {
	return &authHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h *authHandler) Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func (h *authHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBind(&user); err != nil {
			h.logger.Error("Some thing wrong", err)
			ctx.JSON(httpErrors.ErrorResponse(err))
			return
		}
		if err := h.uc.UserValidator(ctx, &user); err != nil {
			ctx.JSON(
				httpErrors.ErrorResponse(
					httpErrors.NewRestError(
						http.StatusBadRequest,
						"",
						err.Error(),
					),
				),
			)
			return
		}
		if token, err := h.uc.Register(ctx, &user); err != nil {
			h.logger.Error(err)
			ctx.JSON(httpErrors.ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusAccepted, token)
			return
		}

	}
}

func (h *authHandler) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		fmt.Println(exists)
		if !exists {
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"error": "Could not found user",
				},
			)
			return
		}
		userModel, ok := user.(*models.User)
		if !ok {
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"error": "Invalid user type",
				},
			)
			return
		}
		c.JSON(
			http.StatusAccepted,
			userModel,
		)
	}
}

func (h *authHandler) Login() gin.HandlerFunc {
	type Login struct {
		Email    string `json:"email" form:"email" validate:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Invalid user data",
				},
			)
			return
		}
		user := &models.User{
			Email:    login.Email,
			Password: login.Password,
		}
		if err := h.uc.UserValidator(c, user); err != nil {
			c.JSON(
				httpErrors.ErrorResponse(
					httpErrors.NewRestError(
						http.StatusBadRequest,
						err.Error(),
						err.Error(),
					),
				),
			)
			return
		}
		if token, err := h.uc.Login(c, user); err != nil {
			c.JSON(
				httpErrors.ErrorResponse(
					err,
				),
			)
			return
		} else {
			c.JSON(
				http.StatusAccepted,
				token,
			)
		}
		fmt.Println(user)
	}
}
