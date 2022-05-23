package ports

import "context"

//go:generate mockgen -source=rating_service_port.go -destination=../../adapters/network/mock/rating_service_mock.go
type RatingService interface {
	GetRatingByMovieId(ctx context.Context, movieId string) (float32, error)
}
