// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"agnos-assignment/app/controller"
	"agnos-assignment/app/service"

	"github.com/google/wire"
)

var patientCtrlSet = wire.NewSet(controller.PatientContollerInit,
	wire.Bind(new(controller.PatientContollerInterface), new(*controller.PatientContoller)),
)

var patientSvcSet = wire.NewSet(service.PatientServiceInit,
	wire.Bind(new(service.PatientServiceInterface), new(*service.PatientService)),
)

type PatientModule struct {
	PatientCtrl controller.PatientContollerInterface
	PatientSvc  service.PatientServiceInterface
}

func NewPatientModule(
	patientCtrl controller.PatientContollerInterface,
	patientSvc service.PatientServiceInterface,
) *PatientModule {
	return &PatientModule{
		PatientCtrl: patientCtrl,
		PatientSvc:  patientSvc,
	}
}

func PatientModuleInit() *PatientModule {
	wire.Build(NewPatientModule, patientCtrlSet, patientSvcSet)
	return nil
}
