package models

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)