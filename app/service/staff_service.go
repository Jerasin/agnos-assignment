package service

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/model"
	"agnos-assignment/app/pkg"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/request"
	"agnos-assignment/app/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StaffServiceInterface interface {
	CreateStaff(c *gin.Context)
	LoginStaff(c *gin.Context)
}

type StaffService struct {
	StaffRepository repository.StaffRepositoryInterface
}

func StaffServiceInit(StaffRepository repository.StaffRepositoryInterface) *StaffService {
	return &StaffService{
		StaffRepository: StaffRepository,
	}
}

func (s StaffService) CreateStaff(c *gin.Context) {
	defer pkg.PanicHandler(c)

	s.StaffRepository.GetBaseRepo().ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		var request request.StaffRequest
		var hospital model.Hospital
		var gender constant.Gender

		if err = c.ShouldBindJSON(&request); err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			pkg.PanicException(constant.InvalidRequest)
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
		request.Password = string(hash)

		err = s.StaffRepository.GetBaseRepo().FindOne(tx, &hospital, "id = ?", request.Hospital)
		if err != nil {
			pkg.PanicException(constant.DataNotFound)
		}

		if request.Gender == "M" {
			gender = constant.Male
		} else {
			gender = constant.Female
		}

		staff := model.Staff{
			Username:    request.Username,
			Password:    request.Password,
			HospitalID:  uint(request.Hospital),
			FirstNameTh: request.FirstNameTh,
			LastNameTh:  request.LastNameTh,
			FirstNameEn: request.FirstNameEn,
			LastNameEn:  request.LastNameEn,
			Email:       request.Email,
			Age:         request.Age,
			Gender:      gender,
		}

		err = s.StaffRepository.GetBaseRepo().Save(tx, &staff)
		if err != nil {
			pkg.PanicDatabaseException(err, c)
		}

		return nil
	})

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.CreateResponse()))
}

func (s StaffService) LoginStaff(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var err error
	var request request.StaffLoginRequest
	var staff model.Staff

	if err = c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = s.StaffRepository.GetBaseRepo().FindOne(nil, &staff, "username = ?", request.Username)
	if err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.PanicException(constant.DataNotFound)
		}

		pkg.PanicException(constant.InvalidRequest)
	}

	isError := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(request.Password))
	if isError != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	jwt := pkg.NewAuthService()

	token := jwt.GenerateToken(staff.Username, staff.HospitalID)

	response := response.LoginStaffModel{
		Token:        token,
		RefreshToken: jwt.GenerateRefreshToken(staff.Username),
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
