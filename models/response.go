package models

/**
 * Common response struct
 */

type Response interface {
	GetStatus() int
	GetMessage() string
	GetData() interface{}
}

func NewResponse(status int, message string, data interface{}) Response {
	return &innerResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type innerResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response innerResponse) GetStatus() int {
	return response.Status
}

func (response innerResponse) GetMessage() string {
	return response.Message
}

func (response innerResponse) GetData() interface{} {
	return response.Data
}
