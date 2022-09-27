package dto

type (
	PondFeeder struct {
		PondUUID   string `json:"pond_uuid"`
		FeederUUID string `json:"feeder_uuid"`
		Name       string `json:"name"`
		Barcode    string `json:"barcode"`
	}

	PondFeederGRPCResponse struct {
		Status  bool         `json:"status"`
		Message string       `json:"message"`
		Data    []PondFeeder `json:"data"`
	}
)
