package v1

import (
	"github.com/diploma/internal/application/ports"
	"github.com/gin-gonic/gin"
)

type v1Handler struct {
	movieUseCase ports.MovieUseCase
	authUseCase  ports.AuthUseCase
}

func NewV1Handler(movieUseCase ports.MovieUseCase, authUseCase ports.AuthUseCase) *v1Handler {
	return &v1Handler{
		movieUseCase: movieUseCase,
		authUseCase:  authUseCase,
	}
}

func (h *v1Handler) Init(v1 *gin.RouterGroup) {
	movies := v1.Group("/movies")
	{
		movies.GET("/", h.getMovies)
		movies.GET("/:id", h.getMovie)
		movies.POST("/", h.createMovie)
		movies.PATCH("/:id", h.updateMovie)
		movies.DELETE("/:id", h.deleteMovie)
	}
	auth := v1.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
	}
}
