package dbs

import "houseService/internal/model"

type FlatRepository interface {
	FlatCreate(*model.Flat) (*model.Flat, error)
	FlatUpdate(update *model.Flat) error
	GetFlatById(id int) (*model.Flat, error)
	HouseFlatsGet(int, string) ([]*model.Flat, error)
}
