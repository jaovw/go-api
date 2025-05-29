package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/schema"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Param id is required")
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
