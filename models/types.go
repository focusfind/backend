package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Spot struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"size:100;not null"`
	Type        string         `gorm:"size:50;not null"`
	Location    datatypes.JSON `gorm:"type:jsonb;not null"` // GeoJSON Point for latitude/longitude
	Description string         `gorm:"size:255"`
	BusyIndex   int            `gorm:"default:0"`
}
