package service

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientServiceInterface interface {
	Search(c *gin.Context)
}

type PatientService struct{}

func PatientServiceInit() *PatientService {
	return &PatientService{}
}

func (s PatientService) Search(c *gin.Context) {
	defer pkg.PanicHandler(c)

	response := map[string]interface{}{
		"message": "ok",
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
