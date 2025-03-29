package controller

import (
	"agnos-assignment/app/request"
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type PatientContollerInterface interface {
	Search(c *gin.Context)
	SearchDetail(c *gin.Context)
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
// @Param ID  path string true "ID"
//
//	@Success		200	{object}	response.PatientSearchModel
//
// @Security Bearer
//
// @Router /patient/search/{ID} [get]
func (p PatientContoller) Search(c *gin.Context) {
	p.svc.Search(c)
}

// @Summary Search Detail Patient
// @Schemes
// @Description Search Detail Patient
// @Tags Patient
//
// @Param   national_id         query     string        false  "string valid"
// @Param   passport_id         query     string        false  "string valid"
// @Param   first_name_en         query     string        false  "string valid"
// @Param   first_name_th         query     string        false  "string valid"
// @Param   middle_name_en         query     string        false  "string valid"
// @Param   middle_name_th         query     string        false  "string valid"
// @Param   last_name_en         query     string        false  "string valid"
// @Param   last_name_th         query     string        false  "string valid"
// @Param   date_of_birth         query     string        false  "string valid"
// @Param   phone_number         query     string        false  "string valid"
// @Param   email         query     string        false  "string valid"
//
//	@Success		200	{object}	response.PatientSearchModel
//
// @Security Bearer
//
// @Router /patient/search [get]
func (p PatientContoller) SearchDetail(c *gin.Context) {
	query := request.CreatePaginationPatientRequest(c)
	p.svc.SearchDetail(c, query)
}
