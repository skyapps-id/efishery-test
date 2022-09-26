package dto

import "time"

type (
	PondResponse struct {
		UUID      string    `json:"uuid"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
