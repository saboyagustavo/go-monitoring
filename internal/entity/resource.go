package entity

import "github.com/google/uuid"

type Resource struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	HTTPMethod string `json:"method"`
	StatusCode int    `json:"status_code"`
}

func NewResource(url, method string, statusCode int) *Resource {
	return &Resource{
		ID:         uuid.New().String(),
		URL:        url,
		HTTPMethod: method,
		StatusCode: statusCode,
	}
}
