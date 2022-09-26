package repository

import (
	"context"
	"feeder-service/entity"

	"gorm.io/gorm"
)

type (
	FeederRepository interface {
		Fatch(ctx context.Context) ([]entity.Feeder, error)
		FatchByID(ctx context.Context, uuid string) (entity.Feeder, error)
		FatchByBarcode(ctx context.Context, barcode []string) ([]entity.Feeder, error)
	}

	feederRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewFeederRepository(orm *gorm.DB) FeederRepository {
	return &feederRepositoryImpl{orm: orm}
}

func (r *feederRepositoryImpl) Fatch(ctx context.Context) ([]entity.Feeder, error) {
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

func (r *feederRepositoryImpl) FatchByID(ctx context.Context, uuid string) (entity.Feeder, error) {
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

func (r *feederRepositoryImpl) FatchByBarcode(ctx context.Context, barcode []string) ([]entity.Feeder, error) {
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
