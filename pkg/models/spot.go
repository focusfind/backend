package models

import (
	"context"
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c Coordinates) GormDataType() string {
	return "geometry(Point,4326)"
}

func (c Coordinates) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_SetSRID(ST_MakePoint(?, ?), 4326)",
		Vars: []interface{}{c.Longitude, c.Latitude},
	}
}

func (c *Coordinates) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// Skip the first 4 bytes (SRID) and the next byte (geometry type)
		if len(v) < 21 {
			return fmt.Errorf("invalid WKB length")
		}
		// Read X coordinate (longitude)
		c.Longitude = float64(int64(v[5])|(int64(v[6])<<8)|(int64(v[7])<<16)|(int64(v[8])<<24)|
			(int64(v[9])<<32)|(int64(v[10])<<40)|(int64(v[11])<<48)|(int64(v[12])<<56)) / 1e11
		// Read Y coordinate (latitude)
		c.Latitude = float64(int64(v[13])|(int64(v[14])<<8)|(int64(v[15])<<16)|(int64(v[16])<<24)|
			(int64(v[17])<<32)|(int64(v[18])<<40)|(int64(v[19])<<48)|(int64(v[20])<<56)) / 1e11
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type Coordinates", value)
	}
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", c.Longitude, c.Latitude), nil
}

type Spot struct {
	gorm.Model
	Name        string      `json:"name" gorm:"unique"`
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates" gorm:"type:geometry(Point,4326)"`
	Description string      `json:"description"`
	BusyIndex   int         `json:"busy_index"`
}
