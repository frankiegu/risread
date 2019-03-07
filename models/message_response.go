package models

type MessageResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
