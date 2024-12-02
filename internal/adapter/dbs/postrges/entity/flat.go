package entity

import "houseService/internal/model"

type Flat struct {
	Id int `db:"flat_number"`

	HouseId int `db:"house_number"`

	Price int32 `db:"price"`

	Rooms int32 `db:"number_of_rooms"`

	Status model.Status `db:"status"`
}
