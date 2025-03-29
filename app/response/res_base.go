package response

type BaseResponse struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
}

type CreateDataResponse struct {
	BaseResponse
	Message string `json:"message" example:"create success"`
}
