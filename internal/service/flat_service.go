package service

import "houseService/internal/model"

//go:generate mockgen -source=flat_service.go -destination=mocks/mock_flat_service.go

type Flat interface {
	FlatCreate(*model.Flat) (*model.Flat, error)
	FlatUpdate(*model.FlatUpdate) (*model.Flat, error)
	HouseFlatsGet(int, string) ([]*model.Flat, error)
}
