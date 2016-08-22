package models

import (
  "time"
)

type ChangeAlert struct {
	AlertId    int64     `json:"alert_id" db:"alert_id"`
	CompareKey string    `json:"compare_key" db:"compare_key"`
	IgnoreNull bool      `json:"ignore_null" db:"ignore_null"`
	QueryKey   string    `json:"query_key" db:"query_key"`
	Timeframe  time.Time `json:"timeframe" db:"timeframe"`
}
