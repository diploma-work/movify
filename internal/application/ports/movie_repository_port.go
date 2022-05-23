package ports

import (
	"context"
	"github.com/diploma/internal/domain"
)

//go:generate mockgen -source=movie_repository_port.go -destination=../../adapters/repository/mock/movie_repository_mock.go
type MovieRepository interface {
	FindById(ctx context.Context, id string) (domain.MovieEntity, error)
	FindAll(ctx context.Context) ([]domain.MovieEntity, error)
	Create(ctx context.Context, movie *domain.MovieEntity) error
	Update(ctx context.Context, movie *domain.MovieEntity) error
	Delete(ctx context.Context, id string) error
}
