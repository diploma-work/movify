package http

import (
	_ "github.com/diploma/docs"
	"github.com/diploma/internal/adapters/handlers/http/v1"
	"github.com/diploma/internal/application/ports"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type handler struct {
	movieUseCase ports.MovieUseCase
	authUseCase  ports.AuthUseCase
}

func NewHandler(movieUseCase ports.MovieUseCase, authUseCase ports.AuthUseCase) *handler {
	return &handler{
		movieUseCase: movieUseCase,
		authUseCase:  authUseCase,
	}
}

func (h *handler) Init() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Handler := v1.NewV1Handler(h.movieUseCase, h.authUseCase)
	api := router.Group("/api")
	v1Routes := api.Group("/v1")
	{
		v1Handler.Init(v1Routes)
	}
	return router
}
