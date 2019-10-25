package entities

import (
	"time"
)

// https://travix.io/encapsulating-dependencies-in-go-b0fd74021f5a
type RefreshToken struct {
	Email     string
	Token     string
	CreatedAt time.Time
	RevokedAt *time.Time
}

func NewRefreshToken(email string, token string) (*RefreshToken, error) {
	return &RefreshToken{
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
