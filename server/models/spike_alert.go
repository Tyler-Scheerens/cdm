package models

import (
  "time"
)

type SpikeAlert struct {
	AlertId        int64     `json:"alert_id"`
	QueryKey       string    `json:"query_key"`
	Timeframe      time.Time `json:"timeframe"`
	UseCountQuery  bool      `json:"use_count_query"`
	UseTermsQuery  bool      `json:"use_terms_query"`
	DocType        string    `json:"doc_type"`
	TermsSize      int32     `json:"terms_size"`
	SpikeHeight    int32     `json:"spike_height"`
	SpikeType      string    `json:"spike_type"`
	AlertOnNewData bool      `json:"alert_on_new_data"`
	ThresholdRef   int32     `json:"threshold_ref"`
	ThresholdCur   int32     `json:"threshold_cur"`
}
