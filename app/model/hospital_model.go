package model

type Hospital struct {
	BaseModel
	NameTh      string `gorm:"column:name_th;type:varchar(255)"`
	NameEn      string `gorm:"column:name_en;type:varchar(255)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255)"`
	Address     string `gorm:"column:address;type:varchar(255)"`
	Staffs      []Staff
	Patients    []Patient
}
