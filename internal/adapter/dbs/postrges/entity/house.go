package entity

import "time"

type House struct {
	Id int `db:"house_number"`

	Address string `db:"address"`

	Year int32 `db:"year_construction"`

	Developer string `db:"developer"`

	CreatedAt time.Time `db:"created_at"`

	UpdateAt time.Time `db:"last_update"`
}
