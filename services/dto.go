package services

type ResponseDTO struct {
	ResponseCode int         `json:"code"`
	Message      string      `json:"msg"`
	Data         interface{} `json:"data"`
}
