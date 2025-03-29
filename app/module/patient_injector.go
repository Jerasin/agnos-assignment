// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"agnos-assignment/app/controller"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/service"

	"github.com/google/wire"
)

var patientCtrlSet = wire.NewSet(controller.PatientContollerInit,
	wire.Bind(new(controller.PatientContollerInterface), new(*controller.PatientContoller)),
)

var patientSvcSet = wire.NewSet(service.PatientServiceInit,
	wire.Bind(new(service.PatientServiceInterface), new(*service.PatientService)),
)

var patientRepoSet = wire.NewSet(repository.PatientRepositoryInit,
	wire.Bind(new(repository.PatientRepositoryInterface), new(*repository.PatientRepository)),
)

type PatientModule struct {
	PatientRepo repository.PatientRepositoryInterface
	PatientCtrl controller.PatientContollerInterface
	PatientSvc  service.PatientServiceInterface
}

func NewPatientModule(
	patientRepo repository.PatientRepositoryInterface,
	patientCtrl controller.PatientContollerInterface,
	patientSvc service.PatientServiceInterface,
) *PatientModule {
	return &PatientModule{
		PatientRepo: patientRepo,
		PatientCtrl: patientCtrl,
		PatientSvc:  patientSvc,
	}
}

func PatientModuleInit() *PatientModule {
	wire.Build(NewPatientModule, patientCtrlSet, patientSvcSet, patientRepoSet, db, baseRepoSet)
	return nil
}
