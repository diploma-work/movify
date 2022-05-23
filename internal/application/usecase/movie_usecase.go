package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/diploma/internal/adapters/dto"
	"github.com/diploma/internal/application/ports"
	"github.com/diploma/internal/domain"
	"github.com/google/uuid"
	"reflect"
)

type movieUseCase struct {
	movieRepository ports.MovieRepository
	ratingService   ports.RatingService
}

func NewMovieUseCase(movieRepository ports.MovieRepository, ratingService ports.RatingService) *movieUseCase {
	return &movieUseCase{
		movieRepository: movieRepository,
		ratingService:   ratingService,
	}
}

func (m movieUseCase) GetById(ctx context.Context, id string) (dto.MovieDto, error) {
	movieModel, err := m.movieRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrRecordNotFound) {
			return dto.MovieDto{}, fmt.Errorf("movie with id=%s %w", id, domain.ErrRecordNotFound)
		}
		return dto.MovieDto{}, err
	}
	movieRating, err := m.ratingService.GetRatingByMovieId(ctx, movieModel.ID)
	if err != nil {
		return dto.MovieDto{}, err
	}
	return m.modelToDto(movieModel, movieRating), nil
}

func (m movieUseCase) GetAll(ctx context.Context) ([]dto.MovieDto, error) {
	movieModels, err := m.movieRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var movieDtos []dto.MovieDto
	for _, movieModel := range movieModels {
		movieRating, err := m.ratingService.GetRatingByMovieId(ctx, movieModel.ID)
		if err != nil {
			return nil, err
		}
		movieDtos = append(movieDtos, m.modelToDto(movieModel, movieRating))
	}
	return movieDtos, nil
}

func (m movieUseCase) Create(ctx context.Context, movieDto dto.MovieDto) (dto.MovieDto, error) {
	movieModel := m.dtoToModel(movieDto)
	movieModel.ID = uuid.New().String()
	if err := m.movieRepository.Create(ctx, &movieModel); err != nil {
		return dto.MovieDto{}, err
	}
	movieDto.ID = movieModel.ID
	return movieDto, nil
}

func (m movieUseCase) Update(ctx context.Context, movieDto dto.MovieDto, id string) (dto.MovieDto, error) {
	movieModel, err := m.movieRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrRecordNotFound) {
			return dto.MovieDto{}, fmt.Errorf("movie with id=%s %w", id, domain.ErrRecordNotFound)
		}
		return dto.MovieDto{}, err
	}
	m.updateOnlyProvidedFields(&movieModel, &movieDto)
	movieRating, err := m.ratingService.GetRatingByMovieId(ctx, movieModel.ID)
	if err != nil {
		return dto.MovieDto{}, err
	}
	return m.modelToDto(movieModel, movieRating), nil
}

func (m movieUseCase) Delete(ctx context.Context, id string) error {
	return m.movieRepository.Delete(ctx, id)
}

func (m movieUseCase) updateOnlyProvidedFields(movieModel *domain.MovieEntity, movieDto *dto.MovieDto) {
	uv := reflect.ValueOf(movieModel)
	de := reflect.ValueOf(movieDto).Elem()
	for i := 0; i < de.NumField(); i++ {
		fieldName := de.Type().Field(i).Name
		fieldValue := de.Field(i).Interface()
		if !reflect.DeepEqual(fieldValue, reflect.Zero(reflect.TypeOf(fieldValue)).Interface()) && fieldName != "ID" {
			reflect.Indirect(uv).FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}
}

func (m movieUseCase) modelToDto(model domain.MovieEntity, movieRating float32) dto.MovieDto {
	return dto.MovieDto{
		ID:          model.ID,
		Title:       model.Title,
		Overview:    model.Overview,
		ReleaseDate: model.ReleaseDate,
		Image:       model.Image,
		Rating:      movieRating,
		Duration:    model.Duration,
		Budget:      model.Budget,
		Genres:      model.Genres,
	}
}

func (m movieUseCase) dtoToModel(movieDto dto.MovieDto) domain.MovieEntity {
	return domain.MovieEntity{
		Title:       movieDto.Title,
		Overview:    movieDto.Overview,
		ReleaseDate: movieDto.ReleaseDate,
		Image:       movieDto.Image,
		Duration:    movieDto.Duration,
		Budget:      movieDto.Budget,
		Genres:      movieDto.Genres,
	}
}
