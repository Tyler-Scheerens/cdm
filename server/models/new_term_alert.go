package models

import (
  "time"
)

type NewTermAlert struct {
	AlertId              int64     `json:"alert_id"`
	QueryKey             string    `json:"query_key"`
	UseTermsQuery        bool      `json:"use_terms_query"`
	DocType              string    `json:"doc_type"`
	TermsSize            int32     `json:"terms_size"`
	Fields               string    `json:"fields"`
	TermsWindowSize      time.Time `json:"terms_window_size"`
	WindowStepSize       time.Time `json:"window_step_size"`
	AlertOnMissingFields bool      `json:"alert_on_missing_fields"`
}
