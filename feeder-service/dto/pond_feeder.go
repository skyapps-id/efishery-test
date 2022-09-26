package dto

type (
	PondFeederResponse struct {
		PondUUID   string `json:"pond_uuid"`
		FeederUUID string `json:"feeder_uuid"`
		Barcode    string `json:"barcode"`
	}
)
