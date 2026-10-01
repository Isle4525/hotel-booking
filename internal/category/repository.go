package category

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Category, error) {
	var categories []Category
	query := "SELECT id, name, slug, created_at, updated_at FROM categories"
	err := r.db.SelectContext(ctx, &categories, query)

	return categories, err
}

func (r *repository) Create(ctx context.Context, category *Category) error {
	query := "INSERT INTO categories (name, slug) VALUES ($1, $2) RETURNING id, created_at, updated_at"
	return r.db.QueryRowxContext(ctx, query, category.Name, category.Slug).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)

}

func (r *repository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
