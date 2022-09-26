package service

import (
	"context"
	"fmt"
	"iot-service/internal_grpc"
	"iot-service/repository"
)

type (
	PondFeedersService interface {
		Fatch(ctx context.Context, PondUuid string) (*string, error)
	}

	pondFeedersServiceImpl struct {
		internal   internal_grpc.PondFeedersInternalGRPC
		repository repository.FeedLogsRepository
	}
)

func NewPondFeedersService(internal internal_grpc.PondFeedersInternalGRPC, repository repository.FeedLogsRepository) PondFeedersService {
	return &pondFeedersServiceImpl{internal: internal, repository: repository}
}

func (i *pondFeedersServiceImpl) Fatch(ctx context.Context, PondUuid string) (*string, error) {
	pondFeeds, _ := i.internal.FatchPondFeeders(ctx, PondUuid)
	fmt.Println(pondFeeds)
	// feedlogs, + := i.i.repository.
	return &pondFeeds.Message, nil
}
