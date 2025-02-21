package repository

import (
	"context"
	"errors"

	"github.com/caophuoclong/whisper/internal/auth"
	"github.com/caophuoclong/whisper/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

// Delete implements auth.Repository.
func (a *authRepo) Delete(ctx *context.Context, userId uuid.UUID) error {
	panic("unimplemented")
}

// Register implements auth.Repository.
func (a *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	result := a.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// FindByEmail implements auth.Repository.
func (a *authRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	result := a.db.WithContext(ctx).Where(
		"email = ?", email,
	).First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, nil
}

// GetUserById implements auth.Repository.
func (a *authRepo) GetUserById(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	var u models.User
	result := a.db.WithContext(ctx).Where(
		"id = ?", userId.String(),
	).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

// Update implements auth.Repository.
func (a *authRepo) Update(ctx *context.Context, user models.User) (*models.User, error) {
	panic("unimplemented")
}

func NewAuthRepo(db *gorm.DB) auth.Repository {
	return &authRepo{
		db: db,
	}
}
