package repository

import (
	"context"
	"feeder-service/entity"

	"gorm.io/gorm"
)

type (
	PondRepository interface {
		FetchByID(ctx context.Context, id string) (entity.Pond, error)
	}

	pondRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewPondRepository(orm *gorm.DB) PondRepository {
	return &pondRepositoryImpl{orm: orm}
}

func (r *pondRepositoryImpl) FetchByID(ctx context.Context, uuid string) (entity.Pond, error) {
	var (
		pond = entity.Pond{}
		err  error
	)

	result := r.orm.WithContext(ctx).Where("uuid = ?", uuid).First(&pond)
	if result.Error == gorm.ErrRecordNotFound {
		return pond, result.Error
	}

	return pond, err
}
