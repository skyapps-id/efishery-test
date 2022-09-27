package entity

import (
	"time"
)

type FeedLogs struct {
	ID                 int       `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	Barcode            string    `gorm:"column:barcode"`
	ScheduleUUID       string    `gorm:"column:schedule_uuid"`
	Data               string    `gorm:"column:data"`
	DataCount          int       `gorm:"column:data_count"`
	OutputGrCount      float64   `gorm:"column:output_gr_count"`
	TotalOutputGrCount float64   `json:"total_output_gr_count" gorm:"->"`
	CreatedAt          time.Time `gorm:"<-:create;column:created_at;type:datetime;autoCreateTime"`
}
