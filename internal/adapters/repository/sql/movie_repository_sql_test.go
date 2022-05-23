package sql_test

import (
	"github.com/diploma/internal/adapters/repository"
	"github.com/diploma/internal/adapters/repository/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieRepository_Save(t *testing.T) {
	db, teardown := sql.TestDB(t, dsn)
	defer teardown("movie_entities")
	repo := repository.NewMovieRepository(db)
	err := repo.Create(ctx, testMovie)
	assert.NoError(t, err)
}

func TestMovieRepository_FindById(t *testing.T) {
	db, teardown := sql.TestDB(t, dsn)
	defer teardown("movie_entities")

	repo := repository.NewMovieRepository(db)
	repo.Create(ctx, testMovie)
	movie, err := repo.FindById(ctx, testMovie.ID)

	assert.NoError(t, err)
	assert.NotNil(t, movie)
}

func TestMovieRepository_Delete(t *testing.T) {
	db, teardown := sql.TestDB(t, dsn)
	defer teardown("movie_entities")

	repo := repository.NewMovieRepository(db)
	repo.Create(ctx, testMovie)
	err := repo.Delete(ctx, testMovie.ID)
	assert.NoError(t, err)
}
