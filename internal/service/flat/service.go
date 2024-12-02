package flat

import (
	"fmt"
	"houseService/internal/adapter/dbs"
	"houseService/internal/model"
	"houseService/internal/service"
	"time"
)

type Service struct {
	repository   dbs.FlatRepository
	houseService service.House
}

func NewService(r dbs.FlatRepository, hs service.House) *Service {
	return &Service{
		repository:   r,
		houseService: hs,
	}
}

func (s *Service) FlatCreate(flat *model.Flat) (*model.Flat, error) {
	flat.Status = model.CREATED

	if err := flat.Validate(); err != nil {
		return nil, fmt.Errorf(`validate flat %w`, service.Invalid)
	}

	newFlat, err := s.repository.FlatCreate(flat)
	if err != nil {
		return nil, err
	}

	house, err := s.houseService.HouseGetById(newFlat.HouseId)
	house.UpdateAt = time.Now()

	err = s.houseService.HouseUpdate(house)
	if err != nil {
		return nil, err
	}

	return newFlat, nil
}

func (s *Service) FlatUpdate(f *model.FlatUpdate) (*model.Flat, error) {
	err := f.Validate()
	if err != nil {
		return nil, fmt.Errorf(`validate flat %w`, service.Invalid)
	}

	flat, err := s.repository.GetFlatById(f.Id)
	if err != nil {
		return nil, err
	}

	if err = flat.SetStatus(f.Status); err != nil {
		return nil, fmt.Errorf(`set status %w`, err)
	}

	if err = s.repository.FlatUpdate(flat); err != nil {
		return nil, err
	}

	return flat, nil
}

func (s *Service) HouseFlatsGet(i int, userType string) ([]*model.Flat, error) {
	var status string
	if userType == string(model.CLIENT) {
		status = string(model.APPROVED)
	}
	flats, err := s.repository.HouseFlatsGet(i, status)
	if err != nil {
		return nil, err
	}
	return flats, nil
}
