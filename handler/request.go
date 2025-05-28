package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateOpeningRequest struct {
	Role     string `json:"role" validate:"required"`
	Company  string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Remote   *bool  `json:"remote" validate:"required"`
	Link     string `json:"link" validate:"required,url"`
	Salary   int64  `json:"salary" validate:"required,gt=0"`
}

// [joaovictor - 27/05/2025] bind the request body and validate
func (r *CreateOpeningRequest) Validate(ctx *gin.Context) bool {
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(r); err != nil {
		logger.Errorf("Error decoding JSON: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid JSON or contains unknown fields")
		return false
	}

	if err := validate.Struct(r); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			first := validationErrors[0]
			msg := "Field '" + first.Field() + "' (type: " + first.Type().String() + ") invalid | rule: " + first.Tag()
			logger.Errorf("Validation error: %s", msg)
			sendError(ctx, http.StatusBadRequest, msg)
			return false
		}

		logger.Errorf("Unknown validation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Validation error")
		return false
	}

	return true
}
