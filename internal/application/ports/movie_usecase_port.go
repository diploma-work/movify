package ports

import (
	"context"
	"github.com/diploma/internal/adapters/dto"
)

type MovieUseCase interface {
	GetById(ctx context.Context, id string) (dto.MovieDto, error)
	GetAll(ctx context.Context) ([]dto.MovieDto, error)
	Create(ctx context.Context, movieDto dto.MovieDto) (dto.MovieDto, error)
	Update(ctx context.Context, movieDto dto.MovieDto, id string) (dto.MovieDto, error)
	Delete(ctx context.Context, id string) error
}
