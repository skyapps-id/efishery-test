package service

import (
	"context"
	"iot-service/dto"
	"iot-service/entity"
	"iot-service/repository"
)

type (
	FeedLogsService interface {
		Create(ctx context.Context, req dto.FeederMessagingRequest) (*dto.FeederMessagingResponse, error)
	}

	feedLogsServiceImpl struct {
		repository repository.FeedLogsRepository
	}
)

func NewFeedLogsService(repository repository.FeedLogsRepository) FeedLogsService {
	return &feedLogsServiceImpl{repository: repository}
}

func (i *feedLogsServiceImpl) Create(ctx context.Context, req dto.FeederMessagingRequest) (*dto.FeederMessagingResponse, error) {
	feedLogs, err := i.repository.Create(ctx, entity.FeedLogs{
		Barcode:       req.Barcode,
		ScheduleUUID:  req.ScheduleUUID,
		Data:          req.Data,
		DataCount:     req.DataCount,
		OutputGrCount: req.OutputGrCount,
	})
	if err != nil {
		return nil, err
	}

	return &dto.FeederMessagingResponse{Barcode: feedLogs.Barcode, ScheduleUUID: feedLogs.ScheduleUUID, Data: feedLogs.Data}, nil
}
