package controller

import (
	"agnos-assignment/app/request"
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type PatientContollerInterface interface {
	Search(c *gin.Context)
	SearchDetail(c *gin.Context)
	GetList(c *gin.Context)
}

type PatientContoller struct {
	*BaseController[service.PatientServiceInterface]
}

func PatientContollerInit(patientService service.PatientServiceInterface) *PatientContoller {
	return &PatientContoller{
		BaseController: BaseControllerInit[service.PatientServiceInterface](patientService),
	}
}

func (p *PatientContoller) GetList(c *gin.Context) {
	p.Svc.GetList(c)
}

func (p *PatientContoller) Search(c *gin.Context) {
	p.Svc.Search(c)
}

func (p *PatientContoller) SearchDetail(c *gin.Context) {
	query := request.CreatePaginationPatientRequest(c)
	p.Svc.SearchDetail(c, query)
}
