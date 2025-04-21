package controller

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/pkg"

	"github.com/gin-gonic/gin"
)

type BaseControllerInterface interface {
	PaginationRoot(c *gin.Context)
}

type BaseController[T any] struct {
	Svc T
}

func BaseControllerInit[T any](svc T) *BaseController[T] {
	return &BaseController[T]{
		Svc: svc,
	}
}

func (b *BaseController[T]) BasePagination(c *gin.Context) {
	defer pkg.PanicHandler(c)

	if svc, ok := any(b.Svc).(interface {
		GetList(c *gin.Context)
	}); ok {
		svc.GetList(c)
	} else {
		pkg.PanicException(constant.MethodNotFound)
	}
}
