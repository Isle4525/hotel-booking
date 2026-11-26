package service

import (
	"context"
	"errors"

	"github.com/Isle4525/hotel-booking/internal/user/dto"
	"github.com/Isle4525/hotel-booking/internal/user/model"
	"github.com/Isle4525/hotel-booking/internal/user/repository"
	"github.com/Isle4525/hotel-booking/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type AuthService struct {
	repo       UserRepo
	jwtManager *jwt.Manager
}

func NewAuthService(repo UserRepo, jwtManager *jwt.Manager) *AuthService {
	return &AuthService{repo: repo, jwtManager: jwtManager}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	return s.repo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", ErrInvalidCredentials
	}

	return s.jwtManager.Generate(user.ID)
}
