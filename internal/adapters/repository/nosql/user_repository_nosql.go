package nosql

import (
	"context"
	"github.com/diploma/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	db *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Collection) *mongoUserRepository {
	return &mongoUserRepository{db}
}

func (m mongoUserRepository) FindByEmail(ctx context.Context, email string) (domain.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoUserRepository) Create(ctx context.Context, userModel *domain.UserEntity) error {
	//TODO implement me
	panic("implement me")
}
