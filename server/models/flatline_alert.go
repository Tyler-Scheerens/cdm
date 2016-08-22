package models

import (
  "time"
)

type FlatlineAlert struct {
	AlertId       int64     `json:"alert_id" db:"alert_id"`
	QueryKey      string    `json:"query_key" db:"query_key"`
	Timeframe     time.Time `json:"timeframe" db:"timeframe"`
	UseCountQuery bool      `json:"use_count_query" db:"use_count_query"`
	DocType       string    `json:"doc_type" db:"doc_type"`
	Threshold     int32     `json:"threshold" db:"threshold"`
}
