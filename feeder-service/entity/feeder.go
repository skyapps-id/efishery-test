package entity

import (
	"time"
)

type Feeder struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	UUID      string    `gorm:"column:uuid"`
	Barcode   string    `gorm:"column:barcode"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"<-:create;column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`
}
