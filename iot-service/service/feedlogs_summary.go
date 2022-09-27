package service

import (
	"context"
	"iot-service/dto"
	"iot-service/internal_const"
	"iot-service/internal_grpc"
	"iot-service/repository"
	"iot-service/server/rest/mapper"
)

type (
	PondFeedersService interface {
		Fetch(ctx context.Context, PondUuid, Date string) (*dto.FeedLogsResponse, error)
	}

	pondFeedersServiceImpl struct {
		internal   internal_grpc.PondFeedersInternalGRPC
		repository repository.FeedLogsRepository
	}
)

func NewPondFeedersService(internal internal_grpc.PondFeedersInternalGRPC, repository repository.FeedLogsRepository) PondFeedersService {
	return &pondFeedersServiceImpl{internal: internal, repository: repository}
}

func (i *pondFeedersServiceImpl) Fetch(ctx context.Context, PondUuid, Date string) (*dto.FeedLogsResponse, error) {
	pondFeeds, err := i.internal.FetchPondFeeders(ctx, PondUuid)
	if err != nil {
		if err.Error() == "not found record" {
			return nil, internal_const.ErrRecordNotFound()
		}
		return nil, internal_const.ErrBadRequest(err)
	}

	var barcodes []string
	for _, item := range pondFeeds.Data {
		barcodes = append(barcodes, item.Barcode)
	}

	feedlogs, err := i.repository.Search(ctx, dto.FeedLogsRequest{Date: Date, Barcode: barcodes})
	if err != nil {
		return nil, err
	}

	return mapper.MapPondFeedersToPondFeedersResponse(Date, pondFeeds, feedlogs), nil
}
