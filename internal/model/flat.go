package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Status string

const (
	CREATED       Status = "created"
	APPROVED      Status = "approved"
	DECLINED      Status = "declined"
	ON_MODERATION Status = "on moderation"
)

type Flat struct {
	Id int

	HouseId int

	Price int32

	Rooms int32

	Status Status
}

type FlatUpdate struct {
	Id int

	Status Status
}

func (f *Flat) Validate() error {
	return validation.ValidateStruct(
		f,
		validation.Field(&f.HouseId, validation.Required, validation.Min(1)),
		validation.Field(&f.Price, validation.Required, validation.Min(0)),
		validation.Field(&f.Rooms, validation.Required, validation.Min(1)),
	)
}

func (f *FlatUpdate) Validate() error {
	return validation.ValidateStruct(
		f,
		validation.Field(&f.Id, validation.Required, validation.Min(1)),
	)
}

func (f *Flat) SetStatus(s Status) error {
	if f.Status == ON_MODERATION && (s == CREATED || s == ON_MODERATION) {
		return fmt.Errorf("flat is already on moderation")
	}

	if f.Status == DECLINED || f.Status == APPROVED {
		return fmt.Errorf("flat has already been moderated")
	}

	f.Status = s
	return nil
}
