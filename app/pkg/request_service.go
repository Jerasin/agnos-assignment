package pkg

import (
	"agnos-assignment/app/constant"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BasePaginationModel struct {
	Page      int
	PageSize  int
	Search    string
	SortField string
	SortValue string
	Offset    int
	Limit     int
}

func CreatePagination(c *gin.Context) *BasePaginationModel {
	reqPage := c.DefaultQuery("page", "1")
	reqPageSize := c.DefaultQuery("pageSize", "15")
	search := c.DefaultQuery("search", "")
	sortField := c.DefaultQuery("sortField", "updated_at")
	sortValue := c.DefaultQuery("sortValue", "desc")

	page, err := strconv.Atoi(reqPage)

	if err != nil {
		log.Error("PaginationModel Convert Data Error: ", err)
		PanicException(constant.ValidateError)
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		log.Error("PaginationModel Convert Data Error: ", err)
		PanicException(constant.ValidateError)
	}

	return &BasePaginationModel{
		Page:      page,
		PageSize:  pageSize,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Offset:    (page - 1) * pageSize,
		Limit:     pageSize,
	}
}
