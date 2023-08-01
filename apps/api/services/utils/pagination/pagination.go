package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParamsFromContext(c *gin.Context) (Params, error) {
	pageQuery, pageQueryOk := c.GetQuery("page")
	pageSizeQuery, pageSizeQueryOk := c.GetQuery("pageSize")

	if !pageQueryOk || !pageSizeQueryOk {
		return Params{}, ErrInvalidPaginationParams
	}

	page, err := strconv.ParseInt(pageQuery, 10, 32)

	if err != nil {
		return Params{}, ErrInvalidPaginationParams
	}

	pageSize, err := strconv.ParseInt(pageSizeQuery, 10, 32)

	if err != nil {
		return Params{}, ErrInvalidPaginationParams
	}

	isPageSizeOk := pageSize > 0 && pageSize <= 90 && pageSize%30 == 0

	if !isPageSizeOk {
		return Params{}, ErrInvalidPaginationParams
	}

	params := Params{
		Page:     int(page),
		PageSize: int(pageSize),
		Offset:   int((page - 1) * pageSize),
	}

	return params, nil
}

func GetPagination(params Params, totalRecords int64) Pagination {
	totalPages := totalRecords / int64(params.PageSize)
	hasPrevious := params.Page > 1
	hasNext := int64(params.Page) < totalPages

	return Pagination{
		Page:         params.Page,
		PageSize:     params.PageSize,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		HasPrevious:  hasPrevious,
		HasNext:      hasNext,
	}
}
