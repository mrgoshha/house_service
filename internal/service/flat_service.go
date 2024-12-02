package service

import "houseService/internal/model"

type Flat interface {
	FlatCreate(*model.Flat) (*model.Flat, error)
	FlatUpdate(*model.FlatUpdate) (*model.Flat, error)
	HouseFlatsGet(int, string) ([]*model.Flat, error)
}
