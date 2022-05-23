package ports

import (
	"context"
	"github.com/diploma/internal/adapters/dto"
)

type AuthUseCase interface {
	Login(ctx context.Context, loginDto dto.LoginDto) (dto.TokenDto, error)
	Register(ctx context.Context, registerDto dto.RegisterDto) (dto.TokenDto, error)
}
