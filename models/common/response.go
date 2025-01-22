package common

import (
	"encoding/json"
	"net/http"
)

/**
 * Common response struct
 */
type response struct {
	Status  int         `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type Response interface {
	GetStatus() int
	GetMessage() string
	GetData() interface{}
}

func New(status int, message string, data interface{}) Response {
	return &response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (response response) GetStatus() int {
	return response.Status
}

func (response response) GetMessage() string {
	return response.Message
}

func (response response) GetData() interface{} {
	return response.Data
}

func SendResponse(rw http.ResponseWriter, status int, message string, data interface{}) {
	response := New(status, message, data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(response)
}
