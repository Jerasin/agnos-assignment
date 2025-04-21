package service

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/model"
	"agnos-assignment/app/pkg"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/request"
	"agnos-assignment/app/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HospitalServiceInterface interface {
	Create(c *gin.Context)
	GetList(c *gin.Context)
	GetDetail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type HospitalService struct {
	BaseSvc            BaseServiceInterface
	HospitalRepository repository.HospitalRepositoryInterface
}

func HospitalServiceInit(baseSvc BaseServiceInterface, HospitalRepository repository.HospitalRepositoryInterface) *HospitalService {
	return &HospitalService{
		BaseSvc:            baseSvc,
		HospitalRepository: HospitalRepository,
	}
}

func (h *HospitalService) Create(c *gin.Context) {
	var body request.Hospital

	err := c.BindJSON(&body)
	if err != nil {
		pkg.PanicException(constant.BadRequest)
	}

	hospital := model.Hospital{
		NameTh:      body.NameTh,
		NameEn:      body.NameEn,
		Address:     body.Address,
		PhoneNumber: body.PhoneNumber,
	}

	conditions := map[string]any{
		"name_th": body.NameTh,
		"name_en": body.NameEn,
	}

	h.BaseSvc.Create(c, conditions, &hospital)
}

func (h *HospitalService) GetList(c *gin.Context) {
	var hospitals []model.Hospital
	var hospital model.Hospital
	var res []response.HospitalModel
	h.BaseSvc.Pagination(c, hospital, &hospitals, &res)
}

func (h *HospitalService) GetDetail(c *gin.Context) {
	var hospital model.Hospital
	var res response.HospitalModel

	h.BaseSvc.GetDetail(c, &hospital, &res)
}

func (h *HospitalService) Update(c *gin.Context) {
	var body request.Hospital
	var hospital model.Hospital

	err := c.BindJSON(&body)

	if err != nil {
		logrus.Infof("err = %+v \n", err)
		pkg.PanicException(constant.BadRequest)
	}

	conditions := map[string]any{
		"name_en": body.NameEn,
		"name_th": body.NameTh,
	}

	h.BaseSvc.IsExist(nil, c, conditions, &hospital)

	h.BaseSvc.Updates(c, &hospital, &body)
}

func (h *HospitalService) Delete(c *gin.Context) {
	var hospital model.Hospital
	h.BaseSvc.Delete(c, &hospital)

}
