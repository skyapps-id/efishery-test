package dto

type (
	DataFeederMessaging struct {
		Timestamp int64   `json:"timestamp"`
		OutputGr  float64 `json:"output_gr"`
	}

	FeederMessaging struct {
		ScheduleUUID string                `json:"schedule_uuid"`
		Data         []DataFeederMessaging `json:"data"`
	}

	FeederMessagingRequest struct {
		Barcode       string  `json:"barcode"`
		ScheduleUUID  string  `json:"schedule_uuid"`
		Data          string  `json:"data"`
		DataCount     int     `json:"data_count"`
		OutputGrCount float64 `json:"output_gr_count"`
	}

	FeederMessagingResponse struct {
		Barcode       string  `json:"barcode"`
		ScheduleUUID  string  `json:"schedule_uuid"`
		Data          string  `json:"data"`
		DataCount     int     `json:"data_count"`
		OutputGrCount float64 `json:"output_gr_count"`
	}
)
