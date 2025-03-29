package request

import "github.com/gin-gonic/gin"

type PatientRequest struct {
	FirstNameTh  string `json:"first_name_th"`
	MiddleNameTh string `json:"middle_name_th"`
	LastNameTh   string `json:"last_name_th"`
	FirstNameEn  string `json:"first_name_en"`
	MiddleNameEn string `json:"middle_name_en"`
	LastNameEn   string `json:"last_name_en"`
	DateOfBirth  string `json:"date_of_birth"`
	NationalId   string `json:"national_id"`
	PassportId   string `json:"passport_id"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
}

type PatientRequestModel struct {
	FirstNameTh  string
	MiddleNameTh string
	LastNameTh   string
	FirstNameEn  string
	MiddleNameEn string
	LastNameEn   string
	DateOfBirth  string
	NationalId   string
	PassportId   string
	PhoneNumber  string
	Email        string
}

func CreatePaginationPatientRequest(c *gin.Context) *PatientRequestModel {
	nationalID := c.DefaultQuery("national_id", "")
	passportID := c.DefaultQuery("passport_id", "")
	firstNameEn := c.DefaultQuery("first_name_en", "")
	firstNameTh := c.DefaultQuery("first_name_th", "")
	middleNameEn := c.DefaultQuery("middle_name_en", "")
	middleNameTh := c.DefaultQuery("middle_name_th", "")
	lastNameEn := c.DefaultQuery("last_name_en", "")
	lastNameTh := c.DefaultQuery("last_name_th", "")
	dateOfBirth := c.DefaultQuery("date_of_birth", "")
	phoneNumber := c.DefaultQuery("phone_number", "")
	email := c.DefaultQuery("email", "")

	return &PatientRequestModel{
		FirstNameTh:  firstNameTh,
		MiddleNameTh: middleNameTh,
		LastNameTh:   lastNameTh,
		FirstNameEn:  firstNameEn,
		MiddleNameEn: middleNameEn,
		LastNameEn:   lastNameEn,
		DateOfBirth:  dateOfBirth,
		NationalId:   nationalID,
		PassportId:   passportID,
		PhoneNumber:  phoneNumber,
		Email:        email,
	}
}
