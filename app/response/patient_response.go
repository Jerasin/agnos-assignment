package response

import "agnos-assignment/app/constant"

type PatientSearchModel struct {
	FirstNameTh  string          `json:"first_name_th"`
	MiddleNameTh string          `json:"middle_name_th"`
	LastNameTh   string          `json:"last_name_th"`
	FirstNameEn  string          `json:"first_name_en"`
	MiddleNameEn string          `json:"middle_name_en"`
	LastNameEn   string          `json:"last_name_en"`
	DateOfBirth  string          `json:"date_of_birth"`
	PatientHn    string          `json:"patient_hn"`
	NationalId   string          `json:"national_id"`
	PassportId   string          `json:"passport_id"`
	PhoneNumber  string          `json:"phone_number"`
	Email        string          `json:"email"`
	Gender       constant.Gender `json:"gender"`
}
