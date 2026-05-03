package validator

import (
	"go-simple-api/internal/delivery/http/dto"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError — ubah error validator jadi list object
func FormatValidationError(err error) []dto.ObjectError {
	errData := []dto.ObjectError{}
	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		errData = append(errData, dto.ObjectError{
			Field:   e.Field(),
			Message: buildErrorMessage(e),
		})
	}

	return errData
}

func buildErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "min":
		return e.Field() + " minimum length is " + e.Param()
	case "max":
		return e.Field() + " maximum length is " + e.Param()
	case "email":
		return e.Field() + " must be a valid email"
	default:
		return e.Field() + " is invalid"
	}
}
