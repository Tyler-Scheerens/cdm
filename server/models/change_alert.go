package models

import (
  "time"
)

type ChangeAlert struct {
	AlertId    int64     `json:"alert_id"`
	CompareKey string    `json:"compare_key"`
	IgnoreNull bool      `json:"ignore_null"`
	QueryKey   string    `json:"query_key"`
	Timeframe  time.Time `json:"timeframe"`
}
