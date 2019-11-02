package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// omitempty to protect against zeroed _id insertion
type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"` //https://stackoverflow.com/questions/55445429/unable-to-decode-the-objectid-subvalue-from-mongodb-results-in-golang
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
