package models

import (
  "time"
)

type CardinalityAlert struct {
	AlertId          int64     `json:"alert_id"`
	QueryKey         string    `json:"query_key"`
	Timeframe        time.Time `json:"timeframe"`
	CardinalityField string    `json:"cardinality_field"`
	MaxCardinality   bool      `json:"max_cardinality"`
	MinCardinality   bool      `json:"min_cardinality"`
}
