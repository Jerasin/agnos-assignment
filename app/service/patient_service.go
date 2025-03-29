package service

import (
	"agnos-assignment/app/constant"

	"agnos-assignment/app/request"

	"agnos-assignment/app/model"
	"agnos-assignment/app/pkg"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientServiceInterface interface {
	Search(c *gin.Context)
	SearchDetail(c *gin.Context, query *request.PatientRequestModel)
}

type PatientService struct {
	PatientRepository repository.PatientRepositoryInterface
}

func PatientServiceInit(patientRepository repository.PatientRepositoryInterface) *PatientService {
	return &PatientService{
		PatientRepository: patientRepository,
	}
}

func (s PatientService) Search(c *gin.Context) {
	defer pkg.PanicHandler(c)
	ID := c.Param("ID")

	var patient model.Patient

	err := s.PatientRepository.GetBaseRepo().FindOne(nil, &patient, "national_id = ? OR passport_id = ?", ID, ID)
	if err != nil {
		pkg.PanicException(constant.DataNotFound)
	}

	var res response.PatientSearchModel
	pkg.ModelDump(&res, patient)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (s PatientService) SearchDetail(c *gin.Context, query *request.PatientRequestModel) {
	defer pkg.PanicHandler(c)

	var patient model.Patient
	payload := pkg.NewAuthService().GetPayloadInToken(c)

	fmt.Println("payload", payload["hospital_id"])

	db := s.PatientRepository.GetBaseRepo().ClientDb().Where("hospital_id = ?", payload["hospital_id"])

	if query.FirstNameTh != "" {
		db.Where("first_name_th = ?", query.FirstNameTh)
	}

	if query.MiddleNameTh != "" {
		db.Where("middle_name_th = ?", query.MiddleNameTh)
	}

	if query.LastNameTh != "" {
		db.Where("last_name_th = ?", query.LastNameTh)
	}

	if query.FirstNameEn != "" {
		db.Where("first_name_en = ?", query.FirstNameEn)
	}

	if query.MiddleNameEn != "" {
		db.Where("middle_name_en = ?", query.MiddleNameEn)
	}

	if query.LastNameEn != "" {
		db.Where("last_name_en = ?", query.LastNameEn)
	}

	if query.NationalId != "" {
		db.Where("national_id = ?", query.NationalId)
	}

	if query.PassportId != "" {
		db.Where("passport_id = ?", query.PassportId)
	}

	if query.PhoneNumber != "" {
		db.Where("phone_number = ?", query.PhoneNumber)
	}

	if query.Email != "" {
		db.Where("email = ?", query.Email)
	}

	if query.DateOfBirth != "" {
		db.Where("date_of_birth = ?", query.DateOfBirth)
	}

	err := db.First(&patient).Error

	if err != nil {
		fmt.Println("err", err)
		pkg.PanicException(constant.DataNotFound)
	}

	var res response.PatientSearchModel
	pkg.ModelDump(&res, patient)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}
