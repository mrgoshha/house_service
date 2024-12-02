package house

import (
	"fmt"
	"houseService/internal/adapter/dbs"
	"houseService/internal/model"
	"houseService/internal/service"
	"time"
)

type Service struct {
	repository dbs.HouseRepository
}

func NewService(r dbs.HouseRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) HouseCreate(req *model.HouseCreate) (*model.House, error) {
	house := &model.House{
		Address:   req.Address,
		Year:      req.Year,
		Developer: req.Developer,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	err := house.Validate()
	if err != nil {
		return nil, fmt.Errorf(`validate house %w`, service.Invalid)
	}

	newHouse, err := s.repository.HouseCreate(house)
	if err != nil {
		return nil, err
	}

	return newHouse, nil
}

func (s *Service) HouseGetById(id int) (*model.House, error) {
	newHouse, err := s.repository.GetHouseById(id)
	if err != nil {
		return nil, err
	}

	return newHouse, nil
}

func (s *Service) HouseUpdate(house *model.House) error {
	err := house.Validate()
	if err != nil {
		return fmt.Errorf(`validate house %w`, service.Invalid)
	}

	err = s.repository.HouseUpdate(house)
	if err != nil {
		return err
	}

	return nil
}
