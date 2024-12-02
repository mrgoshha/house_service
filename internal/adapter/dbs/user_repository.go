package dbs

import "houseService/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByCredentials(user *model.UserLogin) (*model.User, error)
}
