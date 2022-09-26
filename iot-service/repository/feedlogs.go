package repository

import (
	"context"
	"iot-service/dto"
	"iot-service/entity"

	"gorm.io/gorm"
)

type (
	FeedLogsRepository interface {
		Create(ctx context.Context, feedLogs entity.FeedLogs) (entity.FeedLogs, error)
		Search(ctx context.Context, req dto.FeedLogsRequest) ([]entity.FeedLogs, error)
	}

	feedLogsRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewFeedLogsRepository(orm *gorm.DB) FeedLogsRepository {
	return &feedLogsRepositoryImpl{orm: orm}
}

func (r *feedLogsRepositoryImpl) Create(ctx context.Context, feedLogs entity.FeedLogs) (entity.FeedLogs, error) {
	if result := r.orm.WithContext(ctx).Create(&feedLogs); result.Error != nil {
		return feedLogs, result.Error
	}

	return feedLogs, nil
}

func (r *feedLogsRepositoryImpl) Search(ctx context.Context, req dto.FeedLogsRequest) ([]entity.FeedLogs, error) {
	var (
		feedLogs = []entity.FeedLogs{}
		orm      = r.orm.WithContext(ctx)
	)

	if req.Date != nil {
		orm.Where("feed_logs.created_at = ? ", req.Date)
	}

	if req.Barcode != nil {
		orm.Where("feed_logs.barcode IN ?", req.Barcode)
	}

	if result := r.orm.WithContext(ctx).Find(&feedLogs); result.Error != nil {
		return feedLogs, result.Error
	}

	return feedLogs, nil
}
