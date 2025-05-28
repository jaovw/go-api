package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/schema"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if !request.Validate(ctx) {
		return
	}

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
