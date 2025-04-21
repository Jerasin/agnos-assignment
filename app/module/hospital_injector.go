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

var hospitalCtrlSet = wire.NewSet(controller.HospitalContollerInit,
	wire.Bind(new(controller.HospitalContollerInterface), new(*controller.HospitalContoller)),
)

var hospitalSvcSet = wire.NewSet(service.HospitalServiceInit,
	wire.Bind(new(service.HospitalServiceInterface), new(*service.HospitalService)),
)

var hospitalRepoSet = wire.NewSet(repository.HospitalRepositoryInit,
	wire.Bind(new(repository.HospitalRepositoryInterface), new(*repository.HospitalRepository)),
)

type HospitalModule struct {
	HospitalRepo repository.HospitalRepositoryInterface
	HospitalCtrl controller.HospitalContollerInterface
	HospitalSvc  service.HospitalServiceInterface
}

func NewHospitalModule(
	hospitalRepo repository.HospitalRepositoryInterface,
	hospitalCtrl controller.HospitalContollerInterface,
	hospitalSvc service.HospitalServiceInterface,
) *HospitalModule {
	return &HospitalModule{
		HospitalRepo: hospitalRepo,
		HospitalCtrl: hospitalCtrl,
		HospitalSvc:  hospitalSvc,
	}
}

func HospitalModuleInit() *HospitalModule {
	wire.Build(NewHospitalModule, hospitalCtrlSet, hospitalSvcSet, hospitalRepoSet, db, baseRepoSet, baseSvc)
	return nil
}
