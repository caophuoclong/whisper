package auth

import (
	"context"

	"github.com/caophuoclong/whisper/internal/models"
	"github.com/google/uuid"
)

type AuthUsecase interface {
	Register(ctx context.Context, u *models.User) (*models.Token, error)
	UserValidator(ctx context.Context, u *models.User) error
	GetUserById(ctx context.Context, userId uuid.UUID) (*models.User, error)
	Login(ctx context.Context, u *models.User) (*models.Token, error)
}
