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

	// jwtService := c.MustGet("JWTService").(pkg.JWTServiceInterface)
	// payload := jwtService.GetPayloadInToken(c)
	// fmt.Println("payload", payload)

	userContext := c.MustGet("UserContext").(pkg.UserContext)
	fmt.Println("userContext", userContext)

	db := s.PatientRepository.GetBaseRepo().ClientDb()

	db = db.Where("hospital_id = ?", userContext.HospitalID)

	conditions := map[string]string{
		"first_name_th":  query.FirstNameTh,
		"middle_name_th": query.MiddleNameTh,
		"last_name_th":   query.LastNameTh,
		"first_name_en":  query.FirstNameEn,
		"middle_name_en": query.MiddleNameEn,
		"last_name_en":   query.LastNameEn,
		"national_id":    query.NationalId,
		"passport_id":    query.PassportId,
		"phone_number":   query.PhoneNumber,
		"email":          query.Email,
		"date_of_birth":  query.DateOfBirth,
	}

	for field, value := range conditions {
		if value != "" {
			db = db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}
	err := db.First(&patient).Error

	if err != nil {
		fmt.Println("err", err)
		pkg.PanicException(constant.DataNotFound)
	}

	var res response.PatientSearchModel
	pkg.ModelDump(&res, patient)

	fmt.Println("res", res)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}
