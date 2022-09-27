package service

import (
	"context"
	"iot-service/dto"
	"iot-service/internal_grpc"
)

type (
	FeederService interface {
		Fetch(ctx context.Context) ([]dto.FeederResponse, error)
	}

	feederServiceImpl struct {
		internal internal_grpc.FeederInternalGRPC
	}
)

func NewFeederService(internal internal_grpc.FeederInternalGRPC) FeederService {
	return &feederServiceImpl{internal: internal}
}

func (i *feederServiceImpl) Fetch(ctx context.Context) ([]dto.FeederResponse, error) {
	feederGRPC, err := i.internal.FetchFeeder(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.FeederResponse
	for _, row := range feederGRPC.Data {
		result = append(result, dto.FeederResponse{
			UUID:    row.UUID,
			Name:    row.Name,
			Barcode: row.Barcode,
		})
	}

	return result, nil
}
