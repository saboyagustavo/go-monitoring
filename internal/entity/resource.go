package entity

import "github.com/google/uuid"

type Resource struct {
	ID         string
	URL        string
	StatusCode int
}

func NewResource(url string, statusCode int) *Resource {
	return &Resource{
		ID:         uuid.New().String(),
		URL:        url,
		StatusCode: statusCode,
	}
}
