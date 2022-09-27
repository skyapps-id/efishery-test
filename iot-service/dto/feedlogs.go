package dto

type (
	FeedLogsRequest struct {
		Date    string
		Barcode []string
	}

	FeedLogsResponse struct {
		PondUUID string                 `json:"pond_uuid"`
		PondName string                 `json:"pond_name"`
		Feedlogs map[string]interface{} `json:"feedlogs"`
	}
)
