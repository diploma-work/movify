package sql_test

import (
	"context"
	"github.com/diploma/internal/domain"
	"github.com/google/uuid"
	"os"
	"testing"
)

var (
	dsn       string
	testMovie *domain.MovieEntity
	ctx       context.Context
)

func TestMain(m *testing.M) {
	dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	testMovie = &domain.MovieEntity{
		ID:          uuid.New().String(),
		Title:       "Test Movie 1",
		Overview:    "Overview of Test Movie 1",
		ReleaseDate: "01-01-2020",
		Image:       "https://www.ubuy.vn/productimg/?image=aHR0cHM6Ly9tLm1lZGlhLWFtYXpvbi5jb20vaW1hZ2VzL0kvNzFIQk9PN3RZNUwuX0FDX1NMMTUwMF8uanBn.jpg",
		Duration:    150,
		Budget:      100_000_000,
		Genres:      []string{"criminal", "drama"},
	}
	ctx = context.Background()

	os.Exit(m.Run())
}
