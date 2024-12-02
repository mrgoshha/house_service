package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type House struct {
	Id int

	Address string

	Year int32

	Developer string

	CreatedAt time.Time

	UpdateAt time.Time
}

type HouseCreate struct {
	Address string

	Year int32

	Developer string
}

func (f *House) Validate() error {
	return validation.ValidateStruct(
		f,
		validation.Field(&f.Year, validation.Required, validation.Min(1)),
	)
}
