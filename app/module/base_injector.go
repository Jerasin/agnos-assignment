// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"agnos-assignment/app/repository"
	"agnos-assignment/app/utils"

	"github.com/google/wire"
)

var db = wire.NewSet(utils.InitDbClient)

var baseRepoSet = wire.NewSet(repository.BaseRepositoryInit,
	wire.Bind(new(repository.BaseRepositoryInterface), new(*repository.BaseRepository)),
)
