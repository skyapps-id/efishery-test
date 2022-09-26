package dto

import "time"

type (
	FeederResponse struct {
		UUID      string    `json:"uuid"`
		Barcode   string    `json:"barcode"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
