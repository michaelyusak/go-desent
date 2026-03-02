package handler

import (
	"errors"
	"michaelyusak/go-desent.git/apperror"
	"michaelyusak/go-desent.git/entity"
	"michaelyusak/go-desent.git/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	authService service.Auth
}

func NewAuth(
	authService service.Auth,
) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (h *Auth) Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logrus.WithError(err).Warn("[handler][user][Login][ctx.ShouldBindJSON]")

		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	token, err := h.authService.Validate(ctx, user)
	if err != nil {
		var apperr *apperror.AppError
		if errors.As(err, &apperr) {
			ctx.JSON(apperr.Code, map[string]string{
				"message": apperr.Message,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
