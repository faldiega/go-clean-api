package response

import (
	"github.com/labstack/echo/v4"
)

// SendSuccess — kirim success response
func SendSuccess(c echo.Context, httpStatus int, message string, data any) error {
	code := httpStatusToCode(httpStatus)
	return c.JSON(httpStatus, ResponseSuccess(code, message, data))
}

// SendError — kirim error response
func SendError(c echo.Context, httpStatus int, message string, errors any) error {
	code := httpStatusToCode(httpStatus)
	return c.JSON(httpStatus, ResponseError(code, message, errors))
}

// mapping HTTP status → GSA code
func httpStatusToCode(httpStatus int) string {
	switch httpStatus {
	case 200:
		return CodeOK
	case 201:
		return CodeCreated
	case 400:
		return CodeBadRequest
	case 401:
		return CodeUnauthorized
	case 403:
		return CodeForbidden
	case 404:
		return CodeNotFound
	case 408:
		return CodeRequestTimeout
	case 409:
		return CodeConflict
	case 500:
		return CodeInternalServerError
	default:
		return "GSA-" + string(rune(httpStatus))
	}
}
