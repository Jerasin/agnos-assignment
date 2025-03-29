package router

import (
	"agnos-assignment/app/middleware"
	"agnos-assignment/app/module"

	"github.com/gin-gonic/gin"
)

type BaseModuleInit struct {
	PatientModule  *module.PatientModule
	HospitalModule *module.HospitalModule
	StaffModule    *module.StaffModule
}

func NewBaseModule() BaseModuleInit {
	patientInit := module.PatientModuleInit()
	hospitalInit := module.HospitalModuleInit()
	staffInit := module.StaffModuleInit()

	return BaseModuleInit{
		PatientModule:  patientInit,
		HospitalModule: hospitalInit,
		StaffModule:    staffInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	patientRouter := api.Group("patient")
	patientRouter.GET("/search/:ID", init.PatientModule.PatientCtrl.Search)
	patientRouter.Use(middleware.AuthorizeJwt())
	patientRouter.GET("/search", init.PatientModule.PatientCtrl.SearchDetail)

	hospitalRouter := api.Group("hospital")
	hospitalRouter.GET("", init.HospitalModule.HospitalCtrl.GetList)

	staffRouter := api.Group("staff")
	staffRouter.POST("", init.StaffModule.StaffCtrl.CreateStaff)
	staffRouter.POST("/login", init.StaffModule.StaffCtrl.LoginStaff)
	return router
}
