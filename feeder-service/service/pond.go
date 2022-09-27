package service

import (
	"context"
	"feeder-service/dto"
	"feeder-service/repository"
)

type (
	PondService interface {
		FetchByUUID(ctx context.Context, UUID string) (*dto.PondResponse, error)
	}

	pondServiceImpl struct {
		repository repository.PondRepository
	}
)

func NewPondService(repository repository.PondRepository) PondService {
	return &pondServiceImpl{repository: repository}
}

func (i *pondServiceImpl) FetchByUUID(ctx context.Context, UUID string) (*dto.PondResponse, error) {
	pond, err := i.repository.FetchByID(ctx, UUID)
	if err != nil {
		return nil, err
	}

	return &dto.PondResponse{
		UUID:      pond.UUID,
		Name:      pond.Name,
		CreatedAt: pond.CreatedAt,
		UpdatedAt: pond.UpdatedAt,
	}, nil
}
