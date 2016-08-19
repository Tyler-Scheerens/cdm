package models

import (
  "time"
)

type FrequencyAlert struct {
	AlertId       int64     `json:"alert_id"`
	QueryKey      string    `json:"query_key"`
	Timeframe     time.Time `json:"timeframe"`
	NumEvents     int32     `json:"num_events"`
	AttachRelated bool      `json:"attach_related"`
	UseCountQuery bool      `json:"use_count_query"`
	UseTermsQuery bool      `json:"use_terms_query"`
	DocType       string    `json:"doc_type"`
	TermsSize     int32     `json:"terms_size"`
}
