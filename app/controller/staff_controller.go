package controller

import (
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type StaffContollerInterface interface {
	CreateStaff(c *gin.Context)
	LoginStaff(c *gin.Context)
}

type StaffContoller struct {
	svc service.StaffServiceInterface
}

func StaffContollerInit(StaffService service.StaffServiceInterface) *StaffContoller {
	return &StaffContoller{
		svc: StaffService,
	}
}

// @Summary Create Staff
// @Schemes
// @Description Create Staff
// @Tags Staff
//
// @Param request body request.StaffRequest true "query params"
//
//	@Success		200	{object}	response.CreateDataResponse
//
// @Security Bearer
//
// @Router /staff [post]
func (s StaffContoller) CreateStaff(c *gin.Context) {
	s.svc.CreateStaff(c)
}

// @Summary Login Staff
// @Schemes
// @Description Login Staff
// @Tags Staff
//
// @Param request body request.StaffLoginRequest true "query params"
//
//	@Success		200	{object}	response.LoginStaffModel
//
// @Security Bearer
//
// @Router /staff/login [post]
func (s StaffContoller) LoginStaff(c *gin.Context) {
	s.svc.LoginStaff(c)
}
