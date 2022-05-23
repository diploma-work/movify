package network

import (
	"context"
	"math/rand"
	"time"
)

type ratingService struct{}

func NewRatingService() *ratingService {
	return &ratingService{}
}

func (r ratingService) GetRatingByMovieId(ctx context.Context, movieId string) (float32, error) {
	min, max := 0.0, 10.0
	rand.Seed(time.Now().UnixNano())
	rating := min + rand.Float64()*(max-min)
	return float32(rating), nil
}
