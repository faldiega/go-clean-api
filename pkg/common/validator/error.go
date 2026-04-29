package validator

import (
	"github.com/go-playground/validator/v10"
)

// FormatValidationError — ubah error validator jadi map yang readable
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = err.Error()
		return errors
	}

	for _, e := range validationErrors {
		field := e.Field()
		errors[field] = buildErrorMessage(e)
	}

	return errors
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
