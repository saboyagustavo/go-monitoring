package entity

import "github.com/google/uuid"

type Resource struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	HTTPMethod string `json:"http_method"`
	StatusCode int    `json:"status_code"`
}

func NewResource(url, httpMethod string, statusCode int) *Resource {
	return &Resource{
		ID:         uuid.New().String(),
		URL:        url,
		HTTPMethod: httpMethod,
		StatusCode: statusCode,
	}
}
