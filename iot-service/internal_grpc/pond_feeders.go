package internal_grpc

import (
	"context"
	"errors"
	"iot-service/dto"
	"iot-service/pkg"
	service "iot-service/proto/pond_feeders"
)

type (
	PondFeedersInternalGRPC interface {
		FetchPondFeeders(ctx context.Context, pondUuid string) (*dto.PondFeederGRPCResponse, error)
	}

	pondFeedersInternalGRPCImpl struct {
		GRPC pkg.GRPC
	}
)

func NewPondFeedersInternalGRPC(GRPC pkg.GRPC) PondFeedersInternalGRPC {
	return &pondFeedersInternalGRPCImpl{GRPC: GRPC}
}

func (i *pondFeedersInternalGRPCImpl) FetchPondFeeders(ctx context.Context, pondUuid string) (*dto.PondFeederGRPCResponse, error) {
	pondFeeders := service.NewPondFeedersServiceClient(i.GRPC.ClientConn)

	req := &service.PondFeedersRequest{PondUuid: pondUuid}
	response, err := pondFeeders.FetchPondFeeders(ctx, req)
	if err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, errors.New(response.Message)
	}

	if len(response.Data) == 0 {
		return nil, errors.New("not found record")
	}

	var results = []dto.PondFeeder{}
	for _, row := range response.Data {
		results = append(results, dto.PondFeeder{
			PondUUID:   row.PondUuid,
			FeederUUID: row.FeederUuid,
			Name:       row.Name,
			Barcode:    row.Barcode,
		})
	}

	return &dto.PondFeederGRPCResponse{
		Status:  response.Status,
		Message: response.Message,
		Data:    results,
	}, nil
}
