package utils

import (
	"agnos-assignment/app/config"
	"agnos-assignment/app/constant"
	"agnos-assignment/app/model"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InitDataClient struct {
	db *gorm.DB
}

func InitDataClientInit(db *gorm.DB) *InitDataClient {
	return &InitDataClient{
		db: db,
	}
}

func (i InitDataClient) InitPatient() []model.Patient {
	var err error
	var path string
	env := config.GetEnv("APP_ENV", "development")

	if env == "development" {
		path = "app/default_data/patient.json"
	} else {
		path = "default_data/patient.json"
	}

	data := ReadFile(path)

	var newPatientList []model.Patient
	for _, item := range data.([]map[string]any) {
		var gender constant.Gender
		firstNameTh, ok := item["first_name_th"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		firstNameEn, ok := item["first_name_en"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		MiddleNameTh, ok := item["middle_name_th"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		lastNameTh, ok := item["last_name_th"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		MiddleNameEn, ok := item["middle_name_en"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		lastNameEn, ok := item["last_name_en"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		dateOfBirth, ok := item["date_of_birth"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		patientHn, ok := item["patient_hn"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		nationalId, ok := item["national_id"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		passportId, ok := item["passport_id"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		phoneNumber, ok := item["phone_number"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		email, ok := item["email"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		genderBody, ok := item["gender"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		hospitalId, ok := item["hospitalId"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		fmt.Println("hospitalId", hospitalId)
		id, err := strconv.ParseUint(hospitalId, 10, 32)
		if err != nil {
			fmt.Println("err", ok)
			panic(err)
		}
		var uintId uint = uint(id)

		if genderBody == "M" {
			gender = constant.Male
		} else {
			gender = constant.Female
		}

		newPermissionInfo := model.Patient{
			FirstNameTh:  firstNameTh,
			MiddleNameTh: MiddleNameTh,
			LastNameTh:   lastNameTh,
			FirstNameEn:  firstNameEn,
			MiddleNameEn: MiddleNameEn,
			LastNameEn:   lastNameEn,
			DateOfBirth:  dateOfBirth,
			PatientHn:    patientHn,
			NationalId:   nationalId,
			PassportId:   passportId,
			PhoneNumber:  phoneNumber,
			Email:        email,
			Gender:       gender,
			HospitalID:   uintId,
		}

		newPatientList = append(newPatientList, newPermissionInfo)

	}

	var permissionInfoList []model.Patient
	fmt.Printf("permissionInfoList = %v Type = %T \n", permissionInfoList, permissionInfoList)

	err = i.db.Find(&permissionInfoList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(permissionInfoList) == 0 {
		err = i.db.Create(&newPatientList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		fmt.Printf("newPatientList = %v", newPatientList)

		return newPatientList
	} else {
		return permissionInfoList
	}

}

func (i InitDataClient) InitHospital() []model.Hospital {
	var err error
	var path string
	env := config.GetEnv("APP_ENV", "development")

	if env == "development" {
		path = "app/default_data/hospital.json"
	} else {
		path = "default_data/hospital.json"
	}

	data := ReadFile(path)

	var newHospitalList []model.Hospital
	for _, item := range data.([]map[string]any) {
		nameTh, ok := item["name_th"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		nameEn, ok := item["name_en"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		phoneNumber, ok := item["phone_number"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		address, ok := item["address"].(string)

		if !ok {
			fmt.Println("err", ok)
		}

		newHospital := model.Hospital{
			NameTh:      nameTh,
			NameEn:      nameEn,
			PhoneNumber: phoneNumber,
			Address:     address,
		}

		newHospitalList = append(newHospitalList, newHospital)

	}

	var hospitalList []model.Hospital
	fmt.Printf("permissionInfoList = %v Type = %T \n", hospitalList, hospitalList)

	err = i.db.Find(&hospitalList).Error
	if err != nil {
		log.Error(err)
		panic("Find Error")
	}

	if len(hospitalList) == 0 {
		err = i.db.Create(&newHospitalList).Error
		if err != nil {
			log.Error(err)
			panic("Find Error")
		}

		fmt.Printf("newPatientList = %v", newHospitalList)

		return newHospitalList
	} else {
		return hospitalList
	}

}
