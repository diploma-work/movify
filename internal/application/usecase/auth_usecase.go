package usecase

import (
	"context"
	"errors"
	"github.com/diploma/internal/adapters/dto"
	"github.com/diploma/internal/application/ports"
	"github.com/diploma/internal/domain"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authUseCase struct {
	tokenUseCase   ports.TokenUseCase
	userRepository ports.UserRepository
}

func NewAuthUseCase(tokenUseCase ports.TokenUseCase, userRepository ports.UserRepository) *authUseCase {
	return &authUseCase{
		tokenUseCase:   tokenUseCase,
		userRepository: userRepository,
	}
}

func (a authUseCase) Login(ctx context.Context, loginDto dto.LoginDto) (dto.TokenDto, error) {
	userModel, err := a.userRepository.FindByEmail(ctx, loginDto.Email)
	if err != nil {
		return dto.TokenDto{}, domain.ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginDto.Password))
	if err != nil {
		return dto.TokenDto{}, domain.ErrInvalidCredentials
	}
	accessToken, err := a.tokenUseCase.NewJWT(userModel.ID, 30*time.Minute)
	if err != nil {
		return dto.TokenDto{}, err
	}
	return dto.TokenDto{AccessToken: accessToken}, nil
}

func (a authUseCase) Register(ctx context.Context, registerDto dto.RegisterDto) (dto.TokenDto, error) {
	foundUserModel, err := a.userRepository.FindByEmail(ctx, registerDto.Email)
	if err != nil && !errors.Is(err, domain.ErrRecordNotFound) {
		return dto.TokenDto{}, err
	}
	if foundUserModel.ID != "" {
		return dto.TokenDto{}, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), 12)
	userModel := &domain.UserEntity{
		ID:       uuid.New().String(),
		Name:     registerDto.Name,
		Email:    registerDto.Email,
		Password: string(hashedPassword),
		Roles:    pq.StringArray{"USER"},
	}
	if err = a.userRepository.Create(ctx, userModel); err != nil {
		return dto.TokenDto{}, err
	}
	accessToken, err := a.tokenUseCase.NewJWT(userModel.ID, 30*time.Minute)
	if err != nil {
		return dto.TokenDto{}, err
	}
	return dto.TokenDto{AccessToken: accessToken}, nil
}
