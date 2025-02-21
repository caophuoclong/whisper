package auth

import (
	"context"

	"github.com/caophuoclong/whisper/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx *context.Context, user models.User) (*models.User, error)
	Delete(ctx *context.Context, userId uuid.UUID) error
	GetUserById(ctx context.Context, userId uuid.UUID) (*models.User, error)
}
