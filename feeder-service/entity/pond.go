package entity

import (
	"time"
)

type Pond struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	UUID      string    `gorm:"column:uuid"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"<-:create;column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`
}
