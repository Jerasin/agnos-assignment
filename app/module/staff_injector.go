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

var staffCtrlSet = wire.NewSet(controller.StaffContollerInit,
	wire.Bind(new(controller.StaffContollerInterface), new(*controller.StaffContoller)),
)

var staffSvcSet = wire.NewSet(service.StaffServiceInit,
	wire.Bind(new(service.StaffServiceInterface), new(*service.StaffService)),
)

var staffRepoSet = wire.NewSet(repository.StaffRepositoryInit,
	wire.Bind(new(repository.StaffRepositoryInterface), new(*repository.StaffRepository)),
)

type StaffModule struct {
	StaffRepo repository.StaffRepositoryInterface
	StaffCtrl controller.StaffContollerInterface
	StaffSvc  service.StaffServiceInterface
}

func NewStaffModule(
	staffRepo repository.StaffRepositoryInterface,
	staffCtrl controller.StaffContollerInterface,
	staffSvc service.StaffServiceInterface,
) *StaffModule {
	return &StaffModule{
		StaffRepo: staffRepo,
		StaffCtrl: staffCtrl,
		StaffSvc:  staffSvc,
	}
}

func StaffModuleInit() *StaffModule {
	wire.Build(NewStaffModule, staffCtrlSet, staffSvcSet, staffRepoSet, db, baseRepoSet)
	return nil
}
