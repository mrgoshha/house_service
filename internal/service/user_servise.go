package service

import "houseService/internal/model"

type User interface {
	SingUp(user *model.User) (string, error)
	SingIn(user *model.UserLogin) (string, error)
}
