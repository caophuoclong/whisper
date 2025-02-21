package usecase

import (
	"context"
	"net/http"

	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/internal/auth"
	"github.com/caophuoclong/whisper/internal/models"
	"github.com/caophuoclong/whisper/pkg/httpErrors"
	"github.com/caophuoclong/whisper/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type usecase struct {
	authRepo auth.Repository
	cfg      *configs.Config
}

func NewAuthUsecase(authRepo auth.Repository, cfg *configs.Config) auth.AuthUsecase {
	return &usecase{
		authRepo: authRepo,
		cfg:      cfg,
	}
}

func (uc *usecase) Register(ctx context.Context, u *models.User) (*models.Token, error) {
	if existUser, err := uc.authRepo.FindByEmail(ctx, u.Email); existUser != nil || err != nil {
		return nil, httpErrors.NewRestError(
			http.StatusBadRequest,
			"Email already existed",
			"Email already existed",
		)
	}
	if err := u.PrepareCreate(); err != nil {
		return nil, httpErrors.NewRestError(
			http.StatusInternalServerError,
			"Something went wrong!",
			err,
		)
	}
	user, err := uc.authRepo.Register(ctx, u)
	if err != nil {
		return nil, httpErrors.NewRestError(
			http.StatusInternalServerError,
			"Something went wrong!",
			err,
		)
	}

	token := &models.Token{}
	if err := uc.generateToken(user, token); err != nil {
		return nil, err
	}
	return token, nil
}

func (uc *usecase) UserValidator(ctx context.Context, u *models.User) error {
	validate := validator.New()
	err := validate.Struct(u)
	return err
}

func (uc *usecase) GetUserById(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	user, err := uc.authRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *usecase) generateToken(u *models.User, token *models.Token) error {
	if tk, err := utils.GenerateJWTToken(u, uc.cfg, "access"); err == nil {
		token.AccessToken = tk
	} else {
		return httpErrors.NewRestError(
			http.StatusInternalServerError,
			"Failed to generate token",
			err,
		)
	}
	if tk, err := utils.GenerateJWTToken(u, uc.cfg, "refresh"); err == nil {
		token.RefreshToken = tk
	} else {
		return httpErrors.NewRestError(
			http.StatusInternalServerError,
			"Failed to generate token",
			err,
		)
	}
	return nil
}

func (uc *usecase) Login(ctx context.Context, u *models.User) (*models.Token, error) {
	user, err := uc.authRepo.FindByEmail(ctx, u.Email)
	if err != nil || user == nil {
		return nil, httpErrors.NewRestError(
			http.StatusNotFound,
			"Could not found user",
			"Could not found user",
		)
	}

	if err := user.ComparePassword(u.Password); err != nil {
		return nil, httpErrors.NewRestError(
			http.StatusForbidden,
			"Incorrect email or password",
			"Incorrect email or password",
		)
	}
	token := &models.Token{}
	if err := uc.generateToken(user, token); err != nil {
		return nil, err
	}
	return token, nil
}
