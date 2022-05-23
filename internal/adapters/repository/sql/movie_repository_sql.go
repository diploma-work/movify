package sql

import (
	"context"
	"errors"
	"github.com/diploma/internal/domain"
	"gorm.io/gorm"
)

type postgresMovieRepository struct {
	db *gorm.DB
}

func NewPostgresMovieRepository(db *gorm.DB) *postgresMovieRepository {
	return &postgresMovieRepository{db}
}

func (m postgresMovieRepository) FindById(ctx context.Context, id string) (domain.MovieEntity, error) {
	var movie domain.MovieEntity
	if err := m.db.WithContext(ctx).First(&movie, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.MovieEntity{}, domain.ErrRecordNotFound
		}
		return domain.MovieEntity{}, err
	}
	return movie, nil
}

func (m postgresMovieRepository) FindAll(ctx context.Context) ([]domain.MovieEntity, error) {
	var movies []domain.MovieEntity
	if err := m.db.WithContext(ctx).Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (m postgresMovieRepository) Create(ctx context.Context, movie *domain.MovieEntity) error {
	return m.db.WithContext(ctx).Create(&movie).Error
}

func (m postgresMovieRepository) Update(ctx context.Context, movie *domain.MovieEntity) error {
	return m.db.WithContext(ctx).Save(&movie).Error
}

func (m postgresMovieRepository) Delete(ctx context.Context, id string) error {
	return m.db.WithContext(ctx).Delete(&domain.MovieEntity{}, "id = ?", id).Error
}
