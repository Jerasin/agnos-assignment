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
	*BaseController[service.StaffServiceInterface]
}

func StaffContollerInit(StaffService service.StaffServiceInterface) *StaffContoller {
	return &StaffContoller{
		BaseController: BaseControllerInit[service.StaffServiceInterface](StaffService),
	}
}

func (s StaffContoller) CreateStaff(c *gin.Context) {
	s.Svc.CreateStaff(c)
}

func (s StaffContoller) LoginStaff(c *gin.Context) {
	s.Svc.LoginStaff(c)
}
