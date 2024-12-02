package service

import "houseService/internal/model"

type House interface {
	HouseCreate(*model.HouseCreate) (*model.House, error)
	HouseUpdate(*model.House) error
	HouseGetById(int) (*model.House, error)
}
