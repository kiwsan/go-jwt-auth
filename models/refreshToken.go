package models

type RefreshToken interface {
	Token() string
	UserName() string
}
