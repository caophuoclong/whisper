package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"  gorm:"primaryKey"`
	Password  string    `json:"password,omitempty"  form:"password" binding:"required"`
	Email     string    `json:"email"  form:"email" validate:"email" binding:"required" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	LoginDate time.Time `json:"loginDate"`
}

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	u.ID = uuid.New()
	return nil
}

func (u *User) ComparePassword(password string) error {
	fmt.Println(u)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) EmptyPassword() {
	u.Password = ""
}
