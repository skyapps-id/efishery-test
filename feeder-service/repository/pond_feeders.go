package repository

import (
	"context"
	"feeder-service/entity"

	"gorm.io/gorm"
)

type (
	PondFeederRepository interface {
		FetchPondFeeders(ctx context.Context, id string) ([]entity.PondFeeder, error)
	}

	pondFeedersRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewPondFeederRepository(orm *gorm.DB) PondFeederRepository {
	return &pondFeedersRepositoryImpl{orm: orm}
}

func (r *pondFeedersRepositoryImpl) FetchPondFeeders(ctx context.Context, pondUuid string) ([]entity.PondFeeder, error) {
	var (
		pondFeeders = []entity.PondFeeder{}
		err         error
	)

	result := r.orm.WithContext(ctx).Select("pond_feeders.pond_uuid, pond_feeders.feeder_uuid, ponds.name, feeders.barcode").
		Joins("JOIN ponds ponds ON pond_feeders.pond_uuid = ponds.uuid").
		Joins("JOIN feeders feeders ON pond_feeders.feeder_uuid = feeders.uuid").
		Where("pond_feeders.pond_uuid = ?", pondUuid).
		Find(&pondFeeders)
	if result.Error == gorm.ErrRecordNotFound {
		return pondFeeders, result.Error
	}

	return pondFeeders, err
}
