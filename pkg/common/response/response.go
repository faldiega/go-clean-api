package response

// Response code constants
const (
	CodeOK                  = "GCA-200"
	CodeCreated             = "GCA-201"
	CodeBadRequest          = "GCA-400"
	CodeUnauthorized        = "GCA-401"
	CodeForbidden           = "GCA-403"
	CodeNotFound            = "GCA-404"
	CodeRequestTimeout      = "GCA-408"
	CodeConflict            = "GCA-409"
	CodeInternalServerError = "GCA-500"
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
