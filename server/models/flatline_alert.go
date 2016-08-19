package models

import (
  "time"
)

type FlatlineAlert struct {
	AlertId       int64     `json:"alert_id"`
	QueryKey      string    `json:"query_key"`
	Timeframe     time.Time `json:"timeframe"`
	UseCountQuery bool      `json:"use_count_query"`
	DocType       string    `json:"doc_type"`
	Threshold     int32     `json:"threshold"`
}
