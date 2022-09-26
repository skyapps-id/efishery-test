package entity

import (
	"time"
)

type Schedule struct {
	ID            int           `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	UUID          string        `gorm:"column:uuid"`
	PondUUID      string        `gorm:"column:pond_uuid"`
	TimeStart     time.Duration `gorm:"column:time_start"`
	TimeEnd       time.Duration `gorm:"column:time_end"`
	DurationRun   int           `gorm:"column:duration_run"`
	DurationPause int           `gorm:"column:duration_pause"`
	ScheduleType  string        `gorm:"column:schedule_type;type:enum('basic', 'continues', 'advance');default:'basic'"`
	CreatedAt     time.Time     `gorm:"<-:create;column:created_at;type:datetime;autoCreateTime"`
}
