package models

import (
  "time"
)

type CardinalityAlert struct {
	AlertId          int64     `json:"alert_id" db:"alert_id"`
	QueryKey         string    `json:"query_key" db:"query_key`
	Timeframe        time.Time `json:"timeframe" db:"timeframe"`
	CardinalityField string    `json:"cardinality_field" db:"cardinality_field"`
	MaxCardinality   bool      `json:"max_cardinality" db:"max_cardinality"`
	MinCardinality   bool      `json:"min_cardinality" db:"min_cardinality"`
}
