package controller

import (
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type PatientContollerInterface interface {
	Search(c *gin.Context)
}

type PatientContoller struct {
	svc service.PatientServiceInterface
}

func PatientContollerInit(patientService service.PatientServiceInterface) *PatientContoller {
	return &PatientContoller{
		svc: patientService,
	}
}

// @Summary Search Patient
// @Schemes
// @Description Search Patient
// @Tags Patient
//
// @Param   page         query     int        false  "int valid"
// @Param   pageSize         query     int        false  "int valid"
// @Param   sortField         query     string        false  "string valid"
// @Param   sortValue         query     string        false  "string valid"
//
//	@Success		200	{object}	string
//
// @Security Bearer
//
// @Router /users [get]
func (p PatientContoller) Search(c *gin.Context) {
	p.svc.Search(c)
}
