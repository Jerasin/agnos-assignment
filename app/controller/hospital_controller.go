package controller

import (
	"agnos-assignment/app/response"
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type HospitalContollerInterface interface {
	GetList(c *gin.Context)
}

type HospitalContoller struct {
	svc service.HospitalServiceInterface
}

func HospitalContollerInit(HospitalService service.HospitalServiceInterface) *HospitalContoller {
	return &HospitalContoller{
		svc: HospitalService,
	}
}

// @Summary Get List Hospital
// @Schemes
// @Description Search Hospital
// @Tags Hospital
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	response.HospitalModel
//
// @Security Bearer
//
// @Router /hospital [get]
func (h HospitalContoller) GetList(c *gin.Context) {
	query := CreatePagination(c)
	hospital := response.HospitalModel{}
	h.svc.GetPaginationHospital(c, query.page, query.pageSize, query.search, query.sortField, query.sortValue, hospital)
}
