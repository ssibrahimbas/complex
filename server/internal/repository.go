package internal

import (
	"github.com/go-pg/pg/v10"
	"server/entity"
)

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(value *entity.Value) (*entity.Value, error) {
	_, err := r.db.Model(value).Insert()
	return value, err
}

func (r *Repository) GetAll() ([]entity.Value, error) {
	var values []entity.Value
	err := r.db.Model(&values).Select()
	return values, err
}
