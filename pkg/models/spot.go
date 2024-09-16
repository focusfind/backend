package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Spot struct {
	gorm.Model
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Location    datatypes.JSON `json:"location"`
	Description string         `json:"desc"`
	BusyIndex   int            `json:"busy"`

	// 	Location    datatypes.JSON `gorm:"type:jsonb;not null"` // GeoJSON Point for latitude/longitude
}
