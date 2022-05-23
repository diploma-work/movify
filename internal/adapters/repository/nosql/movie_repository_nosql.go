package nosql

import (
	"context"
	"github.com/diploma/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMovieRepository struct {
	db *mongo.Collection
}

func NewMongoMovieRepository(db *mongo.Collection) *mongoMovieRepository {
	return &mongoMovieRepository{db}
}

func (m mongoMovieRepository) FindById(ctx context.Context, id string) (domain.MovieEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoMovieRepository) FindAll(ctx context.Context) ([]domain.MovieEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoMovieRepository) Create(ctx context.Context, movie *domain.MovieEntity) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoMovieRepository) Update(ctx context.Context, movie *domain.MovieEntity) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoMovieRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
