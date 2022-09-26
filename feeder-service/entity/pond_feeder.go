package entity

type PondFeeder struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex"`
	PondUUID   string `gorm:"column:pond_uuid"`
	FeederUUID string `gorm:"column:feeder_uuid"`
	Barcode    string `json:"barcode" gorm:"->"`
}
