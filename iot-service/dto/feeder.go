package dto

type (
	FeederResponse struct {
		UUID    string `json:"pond_uuid"`
		Name    string `json:"name"`
		Barcode string `json:"barcode"`
	}
)
