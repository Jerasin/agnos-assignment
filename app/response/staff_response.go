package response

type (
	StaffModel struct {
		Id       int    `json:"id"`
		Username string `gorm:"unique;not null"`
	}

	LoginStaffModel struct {
		RefreshToken string `json:"refresh_token" binding:"required" example:"admin" validate:"min=1"`
		Token        string `json:"token" binding:"required" example:"1234" validate:"min=1"`
	}
)
