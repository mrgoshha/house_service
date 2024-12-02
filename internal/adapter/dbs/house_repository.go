package dbs

import "houseService/internal/model"

type HouseRepository interface {
	HouseCreate(*model.House) (*model.House, error)
	HouseUpdate(update *model.House) error
	GetHouseById(id int) (*model.House, error)
}
