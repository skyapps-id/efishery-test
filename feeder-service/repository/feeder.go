package repository

import (
	"context"
	"feeder-service/entity"

	"gorm.io/gorm"
)

type (
	FeederRepository interface {
		Fetch(ctx context.Context) ([]entity.Feeder, error)
		FetchByID(ctx context.Context, uuid string) (entity.Feeder, error)
		FetchByBarcode(ctx context.Context, barcode []string) ([]entity.Feeder, error)
	}

	feederRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewFeederRepository(orm *gorm.DB) FeederRepository {
	return &feederRepositoryImpl{orm: orm}
}

func (r *feederRepositoryImpl) Fetch(ctx context.Context) ([]entity.Feeder, error) {
	var (
		feeders = []entity.Feeder{}
		err     error
	)

	result := r.orm.WithContext(ctx).Find(&feeders)
	if result.Error == gorm.ErrRecordNotFound {
		return feeders, result.Error
	}

	return feeders, err
}

func (r *feederRepositoryImpl) FetchByID(ctx context.Context, uuid string) (entity.Feeder, error) {
	var (
		feeder = entity.Feeder{}
		err    error
	)

	result := r.orm.WithContext(ctx).Where("uuid = ?", uuid).First(&feeder)
	if result.Error == gorm.ErrRecordNotFound {
		return feeder, result.Error
	}

	return feeder, err
}

func (r *feederRepositoryImpl) FetchByBarcode(ctx context.Context, barcode []string) ([]entity.Feeder, error) {
	var (
		feeders = []entity.Feeder{}
		err     error
	)

	result := r.orm.WithContext(ctx).Where("barcode IN ?", barcode).Find(&feeders)
	if result.Error == gorm.ErrRecordNotFound {
		return feeders, result.Error
	}

	return feeders, err
}
