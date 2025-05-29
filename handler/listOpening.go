package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/schema"
)

func ListOpeningHandler(ctx *gin.Context) {
	opening := []schema.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing opening")
		return
	}

	sendSuccess(ctx, "list-opening", opening)
}
