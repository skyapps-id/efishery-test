package repository

import (
	"context"
	"iot-service/entity"

	"gorm.io/gorm"
)

type (
	FeedLogsRepository interface {
		Create(ctx context.Context, feedLogs entity.FeedLogs) (entity.FeedLogs, error)
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
