package usecase_test

import (
	"context"
	"github.com/diploma/internal/adapters/dto"
	mockportsnetwork "github.com/diploma/internal/adapters/network/mock"
	mockportsrepository "github.com/diploma/internal/adapters/repository/mock"
	"github.com/diploma/internal/application/usecase"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMovieUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockportsrepository.NewMockMovieRepository(ctrl)
	mockRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil).AnyTimes()
	mockRatingService := mockportsnetwork.NewMockRatingService(ctrl)
	mockRatingService.EXPECT().GetRatingByMovieId(context.Background(), gomock.Any()).Return(float32(5.85), nil).AnyTimes()
	movieUseCase := usecase.NewMovieUseCase(mockRepo, mockRatingService)

	movieDto := dto.MovieDto{
		Title:       "Test movie",
		Overview:    "Test overview",
		ReleaseDate: "01-01-2022",
		Image:       "some-image-url",
		Duration:    120,
		Budget:      1000,
		Genres:      pq.StringArray{"action"},
	}

	res, err := movieUseCase.Create(context.Background(), movieDto)
	require.NoError(t, err)
	require.NotNil(t, res.ID)
}
