package service

import "houseService/internal/model"

//go:generate mockgen -source=house_service.go -destination=mocks/mock_house_service.go

type House interface {
	HouseCreate(*model.HouseCreate) (*model.House, error)
	HouseUpdate(*model.House) error
	HouseGetById(int) (*model.House, error)
}
