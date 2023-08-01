package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrInvalidPaginationParams = errors.New("invalid pagination params")
)

type PaginationParams struct {
	Page     int
	PageSize int
	Offset   int
}

type Pagination struct {
	Page         int   `json:"page"`
	PageSize     int   `json:"pageSize"`
	TotalRecords int64 `json:"totalRecords"`
	TotalPages   int64 `json:"totalPages"`
	HasPrevious  bool  `json:"hasPrevious"`
	HasNext      bool  `json:"hasNext"`
}

func GetPaginationParamsFromContext(c *gin.Context) (PaginationParams, error) {
	pageQuery, pageQueryOk := c.GetQuery("page")
	pageSizeQuery, pageSizeQueryOk := c.GetQuery("pageSize")

	if !pageQueryOk || !pageSizeQueryOk {
		return PaginationParams{}, ErrInvalidPaginationParams
	}

	page, err := strconv.ParseInt(pageQuery, 10, 32)

	if err != nil {
		return PaginationParams{}, ErrInvalidPaginationParams
	}

	pageSize, err := strconv.ParseInt(pageSizeQuery, 10, 32)

	if err != nil {
		return PaginationParams{}, ErrInvalidPaginationParams
	}

	isPageSizeOk := pageSize > 0 && pageSize <= 90 && pageSize%30 == 0

	if !isPageSizeOk {
		return PaginationParams{}, ErrInvalidPaginationParams
	}

	params := PaginationParams{
		Page:     int(page),
		PageSize: int(pageSize),
		Offset:   int((page - 1) * pageSize),
	}

	return params, nil
}

func GetPagination(params PaginationParams, totalRecords int64) Pagination {
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
