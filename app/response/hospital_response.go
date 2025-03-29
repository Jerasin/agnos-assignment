package response

type HospitalModel struct {
	Id          int    `json:"id"`
	NameEn      string `json:"name_en"`
	NameTh      string `json:"name_th"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
