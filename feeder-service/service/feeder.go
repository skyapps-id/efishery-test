package service

import (
	"context"
	"feeder-service/dto"
	"feeder-service/repository"
)

type (
	FeederService interface {
		Fatch(ctx context.Context) (*[]dto.FeederResponse, error)
		FatchByUUID(ctx context.Context, UUID string) (*dto.FeederResponse, error)
		FatchByBarcode(ctx context.Context, Barcode []string) (*[]dto.FeederResponse, error)
	}

	feederServiceImpl struct {
		repository repository.FeederRepository
	}
)

func NewFeederService(repository repository.FeederRepository) FeederService {
	return &feederServiceImpl{repository: repository}
}

func (i *feederServiceImpl) Fatch(ctx context.Context) (*[]dto.FeederResponse, error) {
	feeders, err := i.repository.Fatch(ctx)
	if err != nil {
		return nil, err
	}

	var results []dto.FeederResponse
	for _, row := range feeders {
		results = append(results, dto.FeederResponse{
			UUID:      row.UUID,
			Barcode:   row.Barcode,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return &results, nil
}

func (i *feederServiceImpl) FatchByUUID(ctx context.Context, UUID string) (*dto.FeederResponse, error) {
	feeder, err := i.repository.FatchByID(ctx, UUID)
	if err != nil {
		return nil, err
	}

	return &dto.FeederResponse{
		UUID:      feeder.UUID,
		Barcode:   feeder.Barcode,
		Name:      feeder.Name,
		CreatedAt: feeder.CreatedAt,
		UpdatedAt: feeder.UpdatedAt,
	}, nil
}

func (i *feederServiceImpl) FatchByBarcode(ctx context.Context, Barcode []string) (*[]dto.FeederResponse, error) {
	feeders, err := i.repository.FatchByBarcode(ctx, Barcode)
	if err != nil {
		return nil, err
	}

	var results []dto.FeederResponse
	for _, row := range feeders {
		results = append(results, dto.FeederResponse{
			UUID:      row.UUID,
			Barcode:   row.Barcode,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return &results, nil
}
