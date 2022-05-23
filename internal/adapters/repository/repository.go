package repository

import (
	"github.com/diploma/internal/adapters/repository/sql"
	"github.com/diploma/internal/application/ports"
	"gorm.io/gorm"
)

func NewMovieRepository(db interface{}) ports.MovieRepository {
	return sql.NewPostgresMovieRepository(db.(*gorm.DB))
	//return nosql.NewMongoMovieRepository(db.(*mongo.Collection))
}

func NewUserRepository(db interface{}) ports.UserRepository {
	return sql.NewPostgresUserRepository(db.(*gorm.DB))
	//return nosql.NewMongoUserRepository(db.(*mongo.Collection))
}
