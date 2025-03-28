package router

import (
	"agnos-assignment/app/module"

	"github.com/gin-gonic/gin"
)

type BaseModuleInit struct {
	PatientModule *module.PatientModule
}

func NewBaseModule() BaseModuleInit {
	patientInit := module.PatientModuleInit()
	return BaseModuleInit{
		PatientModule: patientInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	patientRouter := api.Group("patient")
	patientRouter.GET("/search", init.PatientModule.PatientCtrl.Search)

	return router
}
