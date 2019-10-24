package entities

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func IsValidPassword(password string, confirmPassword string) bool {
	return password == confirmPassword
}

func NewUser(username string, password string) (*User, error) {

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:     username,
		PasswordHash: string(encryptPassword),
		Role:         "user",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}, nil
}
