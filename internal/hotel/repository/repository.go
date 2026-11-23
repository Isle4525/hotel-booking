package repository

import (
	"context"

	"github.com/Isle4525/hotel-booking/internal/hotel/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAll(ctx context.Context) ([]model.Hotel, error)
	Create(ctx context.Context, hotel *model.Hotel) error
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]model.Hotel, error) {
	return []model.Hotel{}, nil
}

func (r *repository) Create(ctx context.Context, hotel *model.Hotel) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	return nil
}
