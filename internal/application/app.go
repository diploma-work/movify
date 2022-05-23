package application

import (
	"context"
	"errors"
	transport "github.com/diploma/internal/adapters/handlers/http"
	"github.com/diploma/internal/adapters/network"
	"github.com/diploma/internal/adapters/repository"
	"github.com/diploma/internal/application/usecase"
	"github.com/diploma/internal/config"
	"github.com/diploma/internal/domain"
	"github.com/diploma/internal/infrastructure"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Clean architecture - Movify
// @version 1.0.0
// @description REST-API for movies CRUD with clean architecture

func Start() {
	logger := logrus.New()

	cfg, err := config.Init("configs")
	if err != nil {
		logger.Error("At: config.Init();", err.Error())
		return
	}

	db, err := infrastructure.NewPostgresDatabase(cfg.Postgres)
	if err != nil {
		logger.Error("At: infrastructure.NewPostgresDatabase();", err.Error())
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		logger.Error("At: db.DB();", err.Error())
		return
	}
	if err = db.AutoMigrate(&domain.MovieEntity{}, &domain.UserEntity{}); err != nil {
		logger.Error("At: db.AutoMigrate();", err.Error())
		return
	}

	movieRepository := repository.NewMovieRepository(db)
	userRepository := repository.NewUserRepository(db)
	ratingService := network.NewRatingService()
	movieUseCase := usecase.NewMovieUseCase(movieRepository, ratingService)
	tokenUseCase := usecase.NewJwtTokenUseCase("very-secret-signing-key")
	authUseCase := usecase.NewAuthUseCase(tokenUseCase, userRepository)
	handler := transport.NewHandler(movieUseCase, authUseCase)

	srv := infrastructure.NewServer(cfg.Server, handler.Init())
	go func() {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err = sqlDb.Close(); err != nil {
		logger.Error("At: sqlDb.Close();", err.Error())
	}
}
