package service

import (
	"context"
	"feeder-service/dto"
	"feeder-service/repository"
)

type (
	PondFeederService interface {
		FatchPondFeeders(ctx context.Context, pondUUID string) (*[]dto.PondFeederResponse, error)
	}

	pondFeederServiceImpl struct {
		repository repository.PondFeederRepository
	}
)

func NewPondFeederService(repository repository.PondFeederRepository) PondFeederService {
	return &pondFeederServiceImpl{repository: repository}
}

func (i *pondFeederServiceImpl) FatchPondFeeders(ctx context.Context, pondUUID string) (*[]dto.PondFeederResponse, error) {
	feeders, err := i.repository.FatchPondFeeders(ctx, pondUUID)
	if err != nil {
		return nil, err
	}

	var results []dto.PondFeederResponse
	for _, row := range feeders {
		results = append(results, dto.PondFeederResponse{
			PondUUID:   row.PondUUID,
			FeederUUID: row.FeederUUID,
			Barcode:    row.Barcode,
		})
	}

	return &results, nil
}
