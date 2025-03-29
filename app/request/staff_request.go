package request

import "agnos-assignment/app/constant"

type StaffRequest struct {
	Username    string          `json:"username" binding:"required" example:"test"`
	Password    string          `json:"password" binding:"required" example:"1234"`
	Hospital    int             `json:"hospital" binding:"required" example:"1"`
	FirstNameTh string          `json:"first_name_th" binding:"required" example:"ทดสอบ"`
	LastNameTh  string          `json:"last_name_th" binding:"required" example:"นามสกุล ทดสอบ"`
	FirstNameEn string          `json:"first_name_en" binding:"required" example:"test"`
	LastNameEn  string          `json:"last_name_en" binding:"required" example:"tester"`
	Email       string          `json:"email" binding:"required" example:"admin@gmail.com"`
	Gender      constant.Gender `json:"gender" binding:"required" example:"M"`
	Age         int             `json:"age" binding:"required" example:"10"`
}

type StaffLoginRequest struct {
	Username string `json:"username" binding:"required" example:"test"`
	Password string `json:"password" binding:"required" example:"1234"`
	Hospital int    `json:"hospital" binding:"required" example:"1"`
}
