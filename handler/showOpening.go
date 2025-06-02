package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/schema"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Param id is required")
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
