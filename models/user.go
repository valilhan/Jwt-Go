package models

import (
	"time"
)

type User struct {
	Id           int64     `json:"id" validate:"required"`
	FirstName    *string   `json:"firstName" validate:"required,min=2,max=100"`
	LastName     *string   `json:"lastName" validate:"required,min=2,max=100"`
	Password     *string   `json:"password" validate:"required,min=6,max=100"`
	Email        *string   `json:"email" validate:"required,email,min=6,max=100"`
	Phone        *string   `json:"phone" validate:"required"`
	UserType     *string   `json:"userType" validate:"required"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UserId       string    `json:"userId"`
}
