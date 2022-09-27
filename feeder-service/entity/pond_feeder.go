package entity

type PondFeeder struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	PondUUID   string `gorm:"column:pond_uuid"`
	FeederUUID string `gorm:"column:feeder_uuid"`
	Name       string `json:"name" gorm:"->"`
	Barcode    string `json:"barcode" gorm:"->"`
}
