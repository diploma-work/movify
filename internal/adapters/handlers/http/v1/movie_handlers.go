package v1

import (
	"errors"
	"fmt"
	"github.com/diploma/internal/adapters/dto"
	"github.com/diploma/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Movie By ID
// @Tags movies
// @Description get movie by id
// @Accept json
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} v1.response
// @Failure 400,404 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/movies/{id} [get]
func (h v1Handler) getMovie(c *gin.Context) {
	id := c.Param("id")
	movieDto, err := h.movieUseCase.GetById(c, id)
	if err != nil {
		if errors.As(err, &domain.ErrRecordNotFound) {
			h.respond(c, http.StatusNotFound, nil, err.Error())
			return
		}
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, movieDto, nil)
}

// @Summary Get All Movies
// @Tags movies
// @Description get all movies
// @Accept json
// @Produce json
// @Success 200 {object} v1.response
// @Failure 400 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/movies [get]
func (h v1Handler) getMovies(c *gin.Context) {
	movieDtos, err := h.movieUseCase.GetAll(c)
	if err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, movieDtos, nil)
}

// @Summary Create New Movie
// @Tags movies
// @Description create new movie
// @Accept  json
// @Produce  json
// @Param input body dto.MovieDto true "movie body"
// @Success 200 {object} v1.response
// @Failure 400 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/movies [post]
func (h v1Handler) createMovie(c *gin.Context) {
	var movieDto dto.MovieDto
	if err := c.ShouldBindJSON(&movieDto); err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	movieDto, err := h.movieUseCase.Create(c, movieDto)
	if err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusCreated, movieDto, nil)
}

// @Summary Update Movie
// @Tags movies
// @Description update movie
// @Accept  json
// @Produce  json
// @Param id path string true "movie id"
// @Param input body dto.MovieDto true "movie update body"
// @Success 200 {object} v1.response
// @Failure 400,404 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/movies/{id} [patch]
func (h v1Handler) updateMovie(c *gin.Context) {
	id := c.Param("id")
	var movieDto dto.MovieDto
	if err := c.ShouldBindJSON(&movieDto); err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	movieDto, err := h.movieUseCase.Update(c, movieDto, id)
	if err != nil {
		if errors.As(err, &domain.ErrRecordNotFound) {
			h.respond(c, http.StatusNotFound, nil, err.Error())
			return
		}
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, movieDto, nil)
}

// @Summary Delete Movie
// @Tags movies
// @Description delete movie by id
// @Accept  json
// @Produce  json
// @Param id path string true "movie id"
// @Success 200 {object} v1.response
// @Failure 400,404 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/movies/{id} [delete]
func (h v1Handler) deleteMovie(c *gin.Context) {
	id := c.Param("id")
	if err := h.movieUseCase.Delete(c, id); err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, fmt.Sprintf("Movie with id=%s deleted", id), nil)
}
