package service

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/model"
	"agnos-assignment/app/pkg"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type HospitalServiceInterface interface {
	GetPaginationHospital(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.HospitalModel)
}

type HospitalService struct {
	HospitalRepository repository.HospitalRepositoryInterface
}

func HospitalServiceInit(HospitalRepository repository.HospitalRepositoryInterface) *HospitalService {
	return &HospitalService{
		HospitalRepository: HospitalRepository,
	}
}

func (h HospitalService) GetPaginationHospital(c *gin.Context, page int, pageSize int, search string, sortField string, sortValue string, field response.HospitalModel) {
	defer pkg.PanicHandler(c)

	offset := (page - 1) * pageSize
	limit := pageSize
	fields := DbHandleSelectField(field)

	var hospitals []model.Hospital

	paginationModel := repository.PaginationModel{
		Limit:     limit,
		Offset:    offset,
		Search:    search,
		SortField: sortField,
		SortValue: sortValue,
		Field:     fields,
		Dest:      hospitals,
	}

	data, err := h.HospitalRepository.GetBaseRepo().Pagination(paginationModel, nil)
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	totalPage, err := h.HospitalRepository.GetBaseRepo().TotalPage(&hospitals, pageSize)
	if err != nil {
		log.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	var res []response.HospitalModel
	pkg.ModelDump(&res, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, res, totalPage, page, pageSize))
}
