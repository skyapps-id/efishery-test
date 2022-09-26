package service

import (
	"context"
	"encoding/json"
	"fmt"
	"iot-service/dto"
	"iot-service/internal_grpc"
	"iot-service/repository"
)

type (
	PondFeedersService interface {
		Fatch(ctx context.Context, PondUuid string) (*dto.FeedLogsResponse, error)
	}

	pondFeedersServiceImpl struct {
		internal   internal_grpc.PondFeedersInternalGRPC
		repository repository.FeedLogsRepository
	}
)

func NewPondFeedersService(internal internal_grpc.PondFeedersInternalGRPC, repository repository.FeedLogsRepository) PondFeedersService {
	return &pondFeedersServiceImpl{internal: internal, repository: repository}
}

func (i *pondFeedersServiceImpl) Fatch(ctx context.Context, PondUuid string) (*dto.FeedLogsResponse, error) {
	pondFeeds, err := i.internal.FatchPondFeeders(ctx, PondUuid)
	if err != nil {
		return nil, err
	}
	fmt.Println(pondFeeds)

	feedlogs, _ := i.repository.Search(ctx, dto.FeedLogsRequest{Date: nil, Barcode: nil})
	fmt.Println(feedlogs[0].Data)
	var history map[string]interface{}
	json.Unmarshal([]byte(`{"data": `+feedlogs[0].Data+`}`), &history)
	fmt.Println(history)

	return &dto.FeedLogsResponse{
		PondUUID: pondFeeds.Data[0].PondUUID,
		PondName: "",
		History:  history["data"],
	}, nil
}
