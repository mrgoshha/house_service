package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserType string

const (
	CLIENT    UserType = "client"
	MODERATOR UserType = "moderator"
)

type User struct {
	Id string

	Email string

	Password string

	UserType UserType
}

type UserLogin struct {
	Id string

	Password string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Id, validation.Required, is.UUID),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Length(6, 100)),
		validation.Field(&u.UserType, validation.Required),
	)
}

func (u *UserLogin) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Id, validation.Required, is.UUID),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}
