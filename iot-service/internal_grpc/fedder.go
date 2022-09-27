package internal_grpc

import (
	"context"
	"errors"
	"iot-service/dto"
	"iot-service/pkg"
	service "iot-service/proto/feeder"
)

type (
	FeederInternalGRPC interface {
		FatchFeeder(ctx context.Context) (*dto.FeederGRPCResponse, error)
	}

	feederInternalGRPCImpl struct {
		GRPC pkg.GRPC
	}
)

func NewFeederInternalGRPC(GRPC pkg.GRPC) FeederInternalGRPC {
	return &feederInternalGRPCImpl{GRPC: GRPC}
}

func (i *feederInternalGRPCImpl) FatchFeeder(ctx context.Context) (*dto.FeederGRPCResponse, error) {
	fedders := service.NewFeederServiceClient(i.GRPC.ClientConn)

	req := &service.FeederRequestAll{}
	response, err := fedders.FatchFeeders(ctx, req)
	if err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, errors.New(response.Message)
	}

	if len(response.Data) == 0 {
		return nil, errors.New("not found record")
	}

	var results = []dto.FeederGRPC{}
	for _, row := range response.Data {
		results = append(results, dto.FeederGRPC{
			UUID:    row.Uuid,
			Name:    row.Name,
			Barcode: row.Barcode,
		})
	}

	return &dto.FeederGRPCResponse{
		Status:  response.Status,
		Message: response.Message,
		Data:    results,
	}, nil
}
