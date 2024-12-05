package dbs

import "houseService/internal/model"

//go:generate mockgen -source=flat_repository.go -destination=mocks/mock_flat_repository.go

type FlatRepository interface {
	FlatCreate(*model.Flat) (*model.Flat, error)
	FlatUpdate(update *model.Flat) error
	GetFlatById(id int) (*model.Flat, error)
	HouseFlatsGet(int, string) ([]*model.Flat, error)
}
