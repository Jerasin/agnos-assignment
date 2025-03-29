package model

import "agnos-assignment/app/constant"

type Staff struct {
	BaseModel
	Username    string          `gorm:"column:username;unique;not null"`
	Password    string          `gorm:"column:password;not null"`
	HospitalID  uint            `gorm:"column:hospital_id;not null"`
	FirstNameTh string          `gorm:"column:first_name_th;type:varchar(255)"`
	LastNameTh  string          `gorm:"column:last_name_th;type:varchar(255)"`
	FirstNameEn string          `gorm:"column:first_name_en;type:varchar(255)"`
	LastNameEn  string          `gorm:"column:last_name_en;type:varchar(255)"`
	Email       string          `gorm:"column:email;type:varchar(255)"`
	Gender      constant.Gender `gorm:"column:gender;type:varchar(255)"`
	Age         int             `gorm:"column:age;type:int"`
}
