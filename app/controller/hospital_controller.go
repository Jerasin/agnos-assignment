package controller

import (
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type HospitalContollerInterface interface {
	Create(c *gin.Context)
	GetList(c *gin.Context)
	GetDetail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type HospitalContoller struct {
	*BaseController[service.HospitalServiceInterface]
}

func HospitalContollerInit(hospitalService service.HospitalServiceInterface) *HospitalContoller {
	return &HospitalContoller{
		BaseController: BaseControllerInit[service.HospitalServiceInterface](hospitalService),
	}
}

func (h HospitalContoller) Create(c *gin.Context) {
	h.Svc.Create(c)
}

func (h HospitalContoller) GetList(c *gin.Context) {
	h.BasePagination(c)
}

func (h HospitalContoller) GetDetail(c *gin.Context) {
	h.Svc.GetDetail(c)
}

func (h HospitalContoller) Update(c *gin.Context) {
	h.Svc.Update(c)
}

func (h HospitalContoller) Delete(c *gin.Context) {
	h.Svc.Delete(c)
}
