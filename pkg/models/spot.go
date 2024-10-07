package models

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Spot struct {
	gorm.Model
	Name        string      `json:"name" gorm:"unique"`
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates" gorm:"type:geometry(Point,4326)"`
	Description string      `json:"description"`
	BusyIndex   int         `json:"busy_index"`
}

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

// Scan reads coordinates in the Well-Known Binary (WKB) format
func (c *Coordinates) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		geometry, err := wkb.Unmarshal(v)
		if err != nil {
			return fmt.Errorf("Failed to unmarshal WKB: %w", err)
		}

		// Ensure Point geometry
		point, ok := geometry.(*geom.Point)
		if !ok {
			return fmt.Errorf("Expected Point geometry, got %T", geometry)
		}

		// Extract Longitude and Latitude
		c.Longitude = point.Coords()[0]
		c.Latitude = point.Coords()[1]
		return nil
	default:
		return fmt.Errorf("Unsupported Scan, storing driver.Value type %T into type Coordinates", value)
	}
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", c.Longitude, c.Latitude), nil
}
