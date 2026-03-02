package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(ctx *gin.Context, res any) {
	resMap := map[string]any{
		"message": "ok",
	}

	if res != nil {
		resMap["data"] = res
	}

	ctx.JSON(http.StatusOK, resMap)
}
