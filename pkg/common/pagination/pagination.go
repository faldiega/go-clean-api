package pagination

import (
	"go-simple-api/internal/utils"
	"math"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type MetaData struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type PaginatedResult struct {
	DataItems any      `json:"data_items"`
	MetaData  MetaData `json:"meta_data"`
}

// ambil pagination dari query param
func GetPaginationFromCtx(c echo.Context, defaultPage int, defaultLimit int) Pagination {
	page := utils.ParsePositiveInt(c.QueryParam("page"), defaultPage)
	limit := utils.ParsePositiveInt(c.QueryParam("limit"), defaultLimit)

	// batasi maksimal limit
	if limit > defaultLimit {
		limit = defaultLimit
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

// hitung offset untuk query DB
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

// hitung total pages
func CalculateTotalPages(totalItems, limit int) int {
	return int(math.Ceil(float64(totalItems) / float64(limit)))
}

// build meta response
func BuildMetaData(p Pagination, totalItems int) MetaData {
	return MetaData{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalItems: totalItems,
		TotalPages: CalculateTotalPages(totalItems, p.Limit),
	}
}

// build paginated result
func BuildResult(items any, p Pagination, totalItems int) PaginatedResult {
	return PaginatedResult{
		DataItems: items,
		MetaData:  BuildMetaData(p, totalItems),
	}
}
