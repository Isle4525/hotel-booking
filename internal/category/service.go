package category

import "context"

type Service interface {
	GetAllCategories(ctx context.Context) ([]CategoryRes, error)
	CreateCategory(ctx context.Context, req CategoryReq) (*CategoryRes, error)
	DeleteCategory(ctx context.Context, id int64) error
}

type service struct {
	repo Repository
}

func NewService(reopo Repository) *service {
	return &service{repo: reopo}
}

func (s *service) GetAllCategories(ctx context.Context) ([]CategoryRes, error) {
	cats, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]CategoryRes, len(cats))
	for i, cat := range cats {
		resp[i] = CategoryRes{
			ID:        cat.ID,
			Name:      cat.Name,
			Slug:      cat.Slug,
			CreatedAt: cat.CreatedAt,
			UpdatedAt: cat.UpdatedAt,
		}
	}

	return resp, nil
}

func (s *service) CreateCategory(ctx context.Context, req CategoryReq) (*CategoryRes, error) {
	cat := &Category{
		Name: req.Name,
		Slug: req.Slug,
	}

	if err := s.repo.Create(ctx, cat); err != nil {
		return nil, err
	}

	resp := &CategoryRes{
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
