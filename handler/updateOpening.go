package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/schema"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Param 'id' is required")
		return
	}

	var req UpdateOpeningRequest
	if !req.Validate(ctx) {
		return
	}

	var opening schema.Opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Opening not found with id: %s", id)
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	updates := req.ToMap()

	if len(updates) == 0 {
		sendError(ctx, http.StatusBadRequest, "No fields provided for update")
		return
	}

	if err := db.Model(&opening).Updates(updates).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
