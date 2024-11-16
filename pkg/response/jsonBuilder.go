package response

import (
	"github.com/pndwrzk/cari-barang-service/pkg/constants"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	TotalItems  int `json:"total_items"`
}

type Info struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type Data struct {
	Items      interface{} `json:"items"`
	Pagination Pagination  `json:"pagination,omitempty"`
}

// Response structure for the full JSON response
type Response struct {
	Status string      `json:"status"`
	Info   Info        `json:"info"`
	Data   interface{} `json:"data"`
}

func FetchAll(status string, message string, items interface{}, pagination Pagination) *Response {
	var response Response
	var data Data
	response.Status = status
	response.Info.Message = message
	data.Items = items
	data.Pagination = pagination
	response.Data = data

	return &response
}

func FetchByIdentifier(message string, data interface{}, pagination Pagination) *Response {
	var response Response
	response.Status = constants.SUCCESS_STATUS
	response.Info.Message = message
	response.Data = data

	return &response
}

func ProcessData(status string, message string, data interface{}) *Response {
	var response Response
	response.Status = status
	response.Info.Message = message
	response.Data = data
	return &response
}

func ProcessDataMessageOnly(status string, message string) *Response {
	return ProcessData(status, message, nil)
}

func FailureProcess(status, message string, detail interface{}) *Response {
	var response Response
	response.Status = status
	response.Info.Message = message
	response.Info.Details = detail
	return &response
}
