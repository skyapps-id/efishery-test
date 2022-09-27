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
	)

	subQuery := r.orm.WithContext(ctx).Model(&entity.FeedLogs{}).Select("SUM(feed_logs.output_gr_count)").
		Where("DATE(feed_logs.created_at) = ? ", req.Date).
		Where("feed_logs.barcode IN ?", req.Barcode)
	if result := r.orm.WithContext(ctx).
		Select("*, (?) AS total_output_gr_count", subQuery).
		Where("DATE(feed_logs.created_at) = ? ", req.Date).
		Where("feed_logs.barcode IN ?", req.Barcode).
		Find(&feedLogs); result.Error != nil {
		return feedLogs, result.Error
	}

	return feedLogs, nil
}
