package ports

import (
	"context"
	"github.com/diploma/internal/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (domain.UserEntity, error)
	Create(ctx context.Context, userModel *domain.UserEntity) error
}
