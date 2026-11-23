package service

import (
	"context"

	"github.com/Isle4525/hotel-booking/internal/category/dto"
	"github.com/Isle4525/hotel-booking/internal/category/model"
	"github.com/Isle4525/hotel-booking/internal/category/repository"
)

type Service interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryRes, error)
	CreateCategory(ctx context.Context, req dto.CategoryReq) (*dto.CategoryRes, error)
	DeleteCategory(ctx context.Context, id int64) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllCategories(ctx context.Context) ([]dto.CategoryRes, error) {
	cats, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.CategoryRes, len(cats))
	for i, cat := range cats {
		resp[i] = dto.CategoryRes{
			ID:        cat.ID,
			Name:      cat.Name,
			Slug:      cat.Slug,
			CreatedAt: cat.CreatedAt,
			UpdatedAt: cat.UpdatedAt,
		}
	}

	return resp, nil
}

func (s *service) CreateCategory(ctx context.Context, req dto.CategoryReq) (*dto.CategoryRes, error) {
	cat := &model.Category{
		Name: req.Name,
		Slug: req.Slug,
	}

	if err := s.repo.Create(ctx, cat); err != nil {
		return nil, err
	}

	resp := &dto.CategoryRes{
		ID:        cat.ID,
		Name:      cat.Name,
		Slug:      cat.Slug,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	}

	return resp, nil
}

func (s *service) DeleteCategory(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
