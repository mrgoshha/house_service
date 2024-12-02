package flat

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"houseService/internal/adapter/dbs"
	"houseService/internal/adapter/dbs/postrges"
	"houseService/internal/adapter/dbs/postrges/entity"
	"houseService/internal/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FlatCreate(f *model.Flat) (*model.Flat, error) {
	query := `	INSERT INTO flats (house_number, price, number_of_rooms, status) 
				VALUES ($1, $2, $3, $4) 
				RETURNING flat_number`

	var newFlatId int
	err := r.db.Get(&newFlatId, query,
		f.HouseId, f.Price, f.Rooms, f.Status)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf(`create flat %w`, dbs.ErrorRecordAlreadyExists)
		}
		return nil, fmt.Errorf(`create flat %w`, err)
	}

	f.Id = newFlatId

	return f, nil
}

func (r *Repository) FlatUpdate(f *model.Flat) error {
	query := `	UPDATE flats 
				SET house_number = $1, price = $2, number_of_rooms = $3, status = $4
				WHERE flat_number = $5`

	_, err := r.db.Exec(query, f.HouseId, f.Price, f.Rooms, f.Status, f.Id)

	if err != nil {
		return fmt.Errorf(`update flat %w`, err)
	}

	return nil
}

func (r *Repository) GetFlatById(id int) (*model.Flat, error) {
	query := `	SELECT *
				FROM flats
				WHERE flat_number = $1`

	flat := &entity.Flat{}
	err := r.db.Get(flat, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`get flat by id %w`, dbs.ErrorRecordNotFound)
		}
		return nil, fmt.Errorf(`get flat by id %w`, err)
	}

	return postrges.ToFlatServiceModel(flat), nil
}

func (r *Repository) HouseFlatsGet(houseId int, status string) ([]*model.Flat, error) {
	query := `	SELECT *
	    		FROM flats
				WHERE house_number = $1`

	var args []interface{}
	args = append(args, houseId)

	if status != "" {
		args = append(args, status)
		query += " AND status = $2"
	}

	flats := make([]*entity.Flat, 0)
	err := r.db.Select(&flats, query, args...)

	if err != nil {
		return nil, fmt.Errorf(`get flats in house with number=%d %w`, houseId, err)
	}

	return postrges.ToFlatsServiceModel(flats), nil
}
