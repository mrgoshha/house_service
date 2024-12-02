package entity

import "houseService/internal/model"

type User struct {
	Id string `db:"id"`

	Email string `db:"email"`

	Password string `db:"password"`

	UserType model.UserType `db:"user_type"`
}
