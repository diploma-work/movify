package v1

import (
	"errors"
	"github.com/diploma/internal/adapters/dto"
	"github.com/diploma/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Log In
// @Tags auth
// @Description user log in
// @Accept  json
// @Produce  json
// @Param input body dto.LoginDto true "log in info body"
// @Success 200 {object} v1.response
// @Failure 400 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/auth/login [post]
func (h v1Handler) login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	accessTokenDto, err := h.authUseCase.Login(c, loginDto)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			h.respond(c, http.StatusBadRequest, nil, err.Error())
			return
		}
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, accessTokenDto, nil)
}

// @Summary Registration
// @Tags auth
// @Description user registration
// @Accept  json
// @Produce  json
// @Param input body dto.RegisterDto true "register info body"
// @Success 200 {object} v1.response
// @Failure 400 {object} v1.response
// @Failure 500 {object} v1.response
// @Failure default {object} v1.response
// @Router /api/v1/auth/register [post]
func (h v1Handler) register(c *gin.Context) {
	var registerDto dto.RegisterDto
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	accessTokenDto, err := h.authUseCase.Register(c, registerDto)
	if err != nil {
		h.respond(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	h.respond(c, http.StatusOK, accessTokenDto, nil)
}
