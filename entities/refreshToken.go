package entities

import (
	"time"
)

// https://travix.io/encapsulating-dependencies-in-go-b0fd74021f5a
type RefreshToken struct {
	UserId    string
	Email     string
	Token     string
	CreatedAt time.Time
	RevokedAt *time.Time
}

func NewRefreshToken(userId string, email string, token string) (*RefreshToken, error) {
	return &RefreshToken{
		UserId:    userId,
		Email:     email,
		Token:     token,
		CreatedAt: time.Now().UTC(),
		RevokedAt: &time.Time{},
	}, nil
}

func RevokeRefreshToken(email string, token string) (*RefreshToken, error) {

	revokedAt := time.Now().UTC()

	return &RefreshToken{
		Email:     email,
		Token:     token,
		CreatedAt: time.Now().UTC(),
		RevokedAt: &revokedAt,
	}, nil
}
