package dto

type ObjectError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
