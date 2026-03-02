package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Common struct{}

func (h *Common) Ping(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, map[string]any{
		"success": true,
	})
}

func (h *Common) Echo(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if len(body) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "empty body",
		})
		return
	}

	logrus.WithField("body", string(body)).
		Info("[handler][common][echo] got payload")

	ctx.Data(http.StatusOK, "application/json", body)
}
