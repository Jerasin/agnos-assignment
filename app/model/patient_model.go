package model

type Patient struct {
	BaseModel
	TotalPrice  float64 `gorm:"not null"`
	TotalAmount int     `gorm:"not null"`
	WalletID    uint    `gorm:"not null"`
	CreatedBy   uint    `gorm:"not null"`
}
