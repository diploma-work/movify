package sql

import (
	"context"
	"errors"
	"github.com/diploma/internal/domain"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *postgresUserRepository {
	return &postgresUserRepository{db}
}

func (u postgresUserRepository) FindByEmail(ctx context.Context, email string) (domain.UserEntity, error) {
	var userModel domain.UserEntity
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.UserEntity{}, domain.ErrRecordNotFound
		}
		return domain.UserEntity{}, err
	}
	return userModel, nil
}

func (u postgresUserRepository) Create(ctx context.Context, userModel *domain.UserEntity) error {
	return u.db.WithContext(ctx).Create(&userModel).Error
}
