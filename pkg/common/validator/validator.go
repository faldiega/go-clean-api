package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

// CustomValidator wraps go-playground/validator
type CustomValidator struct {
	validator *validator.Validate
}

// singleton — hanya satu instance validator
var (
	instance *CustomValidator
	once     sync.Once
)

// GetValidator — ambil instance validator (singleton)
func GetValidator() *CustomValidator {
	once.Do(func() {
		instance = &CustomValidator{
			validator: validator.New(),
		}
	})
	return instance
}

// Validate — implement echo.Validator interface
func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}
