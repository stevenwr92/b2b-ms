package dto

type Register struct{}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}