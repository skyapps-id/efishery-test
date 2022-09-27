package dto

type (
	FeederGRPC struct {
		UUID    string `json:"pond_uuid"`
		Name    string `json:"name"`
		Barcode string `json:"barcode"`
	}

	FeederGRPCResponse struct {
		Status  bool         `json:"status"`
		Message string       `json:"message"`
		Data    []FeederGRPC `json:"data"`
	}
)
