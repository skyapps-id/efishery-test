package mapper

import (
	"encoding/json"
	"iot-service/dto"
	"iot-service/entity"
)

type (
	Details struct {
		FeederUUID    string      `json:"feeder_uuid"`
		FeederBarcode string      `json:"feeder_barcode"`
		History       interface{} `json:"history"`
	}
	Content struct {
		TotalOutputGr float64   `json:"total_output_gr"`
		Details       []Details `json:"details"`
	}
)

func MapPondFeedersToPondFeedersResponse(Date string, pondFeeds *dto.PondFeederGRPCResponse, feedlogs []entity.FeedLogs) *dto.FeedLogsResponse {
	var details []Details
	for _, log := range pondFeeds.Data {
		var isAppend = false
		for _, row := range feedlogs {
			if log.Barcode == row.Barcode {
				var history map[string]interface{}
				json.Unmarshal([]byte(row.Data), &history)
				details = append(details, Details{
					FeederUUID:    log.FeederUUID,
					FeederBarcode: log.Barcode,
					History:       history["data"],
				})
				isAppend = true
				break
			}
		}
		if !isAppend {
			details = append(details, Details{
				FeederUUID:    log.FeederUUID,
				FeederBarcode: log.Barcode,
				History:       nil,
			})
		}
	}

	content := Content{
		TotalOutputGr: feedlogs[0].TotalOutputGrCount,
		Details:       details,
	}

	mapFeedlogs := map[string]interface{}{
		Date: content,
	}

	return &dto.FeedLogsResponse{
		PondUUID: pondFeeds.Data[0].PondUUID,
		PondName: pondFeeds.Data[0].Name,
		Feedlogs: mapFeedlogs,
	}
}
