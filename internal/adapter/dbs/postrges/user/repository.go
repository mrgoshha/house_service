package user

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

func (r *Repository) CreateUser(u *model.User) error {
	query := `	INSERT INTO users (id, email, password, user_type) 
				VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query,
		u.Id, u.Email, u.Password, u.UserType)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf(`create user %w`, dbs.ErrorRecordAlreadyExists)
		}
		return fmt.Errorf(`create user %w`, err)
	}

	return nil
}

func (r *Repository) GetUserByCredentials(u *model.UserLogin) (*model.User, error) {
	query := `	SELECT *
				FROM users
				WHERE id = $1 AND password = $2`

	user := &entity.User{}
	err := r.db.Get(user, query, u.Id, u.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`get user by credentials %w`, dbs.ErrorRecordNotFound)
		}
		return nil, fmt.Errorf(`get user by credentials %w`, err)
	}

	return postrges.ToUserServiceModel(user), nil
}
