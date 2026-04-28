package response

// Response code constants
const (
	CodeOK                  = "GSA-200"
	CodeCreated             = "GSA-201"
	CodeBadRequest          = "GSA-400"
	CodeUnauthorized        = "GSA-401"
	CodeForbidden           = "GSA-403"
	CodeNotFound            = "GSA-404"
	CodeConflict            = "GSA-409"
	CodeInternalServerError = "GSA-500"
)

// Base response struct
type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Errors  any    `json:"errors"`
}

// Success — data diisi, errors null
func ResponseSuccess(code string, message string, data any) BaseResponse {
	return BaseResponse{
		Code:    code,
		Data:    data,
		Errors:  nil,
		Message: message,
	}
}

// Error — data null, errors diisi
func ResponseError(code string, message string, errors any) BaseResponse {
	return BaseResponse{
		Code:    code,
		Data:    nil,
		Errors:  errors,
		Message: message,
	}
}
