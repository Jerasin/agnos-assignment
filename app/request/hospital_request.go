package request

type Hospital struct {
	NameTh      string `json:"name_th"`
	NameEn      string `json:"name_en"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
