package model

import "agnos-assignment/app/constant"

type Patient struct {
	BaseModel
	FirstNameTh  string          `gorm:"column:first_name_th;type:varchar(255)"`
	MiddleNameTh string          `gorm:"column:middle_name_th;type:varchar(255)"`
	LastNameTh   string          `gorm:"column:last_name_th;type:varchar(255)"`
	FirstNameEn  string          `gorm:"column:first_name_en;type:varchar(255)"`
	MiddleNameEn string          `gorm:"column:middle_name_en;type:varchar(255)"`
	LastNameEn   string          `gorm:"column:last_name_en;type:varchar(255)"`
	DateOfBirth  string          `gorm:"column:date_of_birth;type:varchar(255)"`
	PatientHn    string          `gorm:"column:patient_hn;type:varchar(255)"`
	NationalId   string          `gorm:"unique;column:national_id;type:varchar(255)"`
	PassportId   string          `gorm:"unique;column:passport_id;type:varchar(255)"`
	PhoneNumber  string          `gorm:"column:phone_number;type:varchar(255)"`
	Email        string          `gorm:"column:email;type:varchar(255)"`
	Gender       constant.Gender `gorm:"column:gender;type:varchar(255)"`
	HospitalID   uint            `gorm:"column:hospital_id;not null"`
}
