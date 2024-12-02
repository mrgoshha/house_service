package house

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

func (r *Repository) HouseCreate(h *model.House) (*model.House, error) {
	query := `	INSERT INTO houses (address, year_construction, developer, created_at, last_update) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING house_number`

	var newHouseId int
	err := r.db.Get(&newHouseId, query,
		h.Address, h.Year, h.Developer, h.CreatedAt, h.UpdateAt)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf(`create house %w`, dbs.ErrorRecordAlreadyExists)
		}
		return nil, fmt.Errorf(`create house %w`, err)
	}

	h.Id = newHouseId

	return h, nil
}

func (r *Repository) HouseUpdate(h *model.House) error {
	query := `	UPDATE houses 
				SET address = $1, year_construction = $2, developer = $3,
				    created_at = $4, last_update = $5
				WHERE house_number = $6`

	_, err := r.db.Exec(query, h.Address, h.Year, h.Developer, h.CreatedAt, h.UpdateAt, h.Id)

	if err != nil {
		return fmt.Errorf(`update house %w`, err)
	}

	return nil
}

func (r *Repository) GetHouseById(id int) (*model.House, error) {
	query := `	SELECT *
				FROM houses
				WHERE house_number = $1`

	house := &entity.House{}
	err := r.db.Get(house, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`get house by id %w`, dbs.ErrorRecordNotFound)
		}
		return nil, fmt.Errorf(`get house by id %w`, err)
	}

	return postrges.ToHouseServiceModel(house), nil
}
