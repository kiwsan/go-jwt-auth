package entities

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func IsValidPassword(password string, confirmPassword string) bool {
	return password == confirmPassword
}

func NewUser(email string, password string) (*User, error) {

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:        email,
		PasswordHash: string(encryptPassword),
		Role:         "user",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}, nil
}
