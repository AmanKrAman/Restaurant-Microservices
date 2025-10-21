package utils

import (
	"encoding/json"
	"time"
)

type Response struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
}

type ResponseBuilder struct{}

func (rb *ResponseBuilder) GenerateResponse(status bool, message string, data interface{}, meta interface{}) (string, error) {

	var responseStatus string
	if status {
		responseStatus = "ok"
	} else {
		responseStatus = "error"
	}

	timestamp := time.Now().Format(time.RFC3339)

	response := Response{
		Status:    responseStatus,
		Message:   message,
		Timestamp: timestamp,
		Data:      data,
		Meta:      meta,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(jsonResponse), nil
}
