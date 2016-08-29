package models

import (
  "time"

  "github.com/jmoiron/sqlx/types"
)

type Alert struct {
	AlertId            int64  `json:"alert_id" db:"alert_id"`
	ElasticsearchId    int64 `json:"elasticsearch_id" db:"elasticsearch_id"`
	RuleName           string `json:"rule_name" db:"rule_name"`
	RuleType           string `json:"rule_type" db:"rule_type"`
  AlertType          string `json:"alert_type" db:"alert_type"`
  ElasticsearchIndex string `json:"elasticsearch_index" db:"elasticsearch_index"`
  AlertOptions       types.JSONText `json:"-" db:"alert_options"`
}

// options specific to each alert type
type AnyAlert struct {
	QueryKey string `json:"query_key" db:"query_key"`
}

type BlacklistAlert struct {
	CompareKey string `json:"compare_key" db:"compare_key"`
	Blacklist  string `json:"blacklist" db:"blacklist"`
}

type CardinalityAlert struct {
	QueryKey         string    `json:"query_key" db:"query_key`
	Timeframe        time.Time `json:"timeframe" db:"timeframe"`
	CardinalityField string    `json:"cardinality_field" db:"cardinality_field"`
	MaxCardinality   bool      `json:"max_cardinality" db:"max_cardinality"`
	MinCardinality   bool      `json:"min_cardinality" db:"min_cardinality"`
}

type ChangeAlert struct {
	CompareKey string    `json:"compare_key" db:"compare_key"`
	IgnoreNull bool      `json:"ignore_null" db:"ignore_null"`
	QueryKey   string    `json:"query_key" db:"query_key"`
	Timeframe  time.Time `json:"timeframe" db:"timeframe"`
}

type FlatlineAlert struct {
	QueryKey      string    `json:"query_key" db:"query_key"`
	Timeframe     time.Time `json:"timeframe" db:"timeframe"`
	UseCountQuery bool      `json:"use_count_query" db:"use_count_query"`
	DocType       string    `json:"doc_type" db:"doc_type"`
	Threshold     int32     `json:"threshold" db:"threshold"`
}

type FrequencyAlert struct {
	QueryKey      string    `json:"query_key" db:"query_key"`
	Timeframe     time.Time `json:"timeframe" db:"timeframe"`
	NumEvents     int32     `json:"num_events" db:"num_events"`
	AttachRelated bool      `json:"attach_related" db:"attach_related"`
	UseCountQuery bool      `json:"use_count_query" db:"use_count_query"`
	UseTermsQuery bool      `json:"use_terms_query" db:"use_terms_query"`
	DocType       string    `json:"doc_type" db:"doc_type"`
	TermsSize     int32     `json:"terms_size" db:"terms_size"`
}

type NewTermAlert struct {
	QueryKey             string    `json:"query_key" db:"query_key"`
	UseTermsQuery        bool      `json:"use_terms_query" db:"use_terms_query"`
	DocType              string    `json:"doc_type" db:"doc_type"`
	TermsSize            int32     `json:"terms_size" db:"terms_size"`
	Fields               string    `json:"fields" db:"fields"`
	TermsWindowSize      time.Time `json:"terms_window_size" db:"terms_window_size"`
	WindowStepSize       time.Time `json:"window_step_size" db:"window_step_size"`
	AlertOnMissingFields bool      `json:"alert_on_missing_fields" db:"alert_on_missing_fields"`
}

type SpikeAlert struct {
	QueryKey       string    `json:"query_key" db:"query_key"`
	Timeframe      time.Time `json:"timeframe" db:"timeframe"`
	UseCountQuery  bool      `json:"use_count_query" db:"use_count_query"`
	UseTermsQuery  bool      `json:"use_terms_query" db:"use_terms_query"`
	DocType        string    `json:"doc_type" db:"doc_type"`
	TermsSize      int32     `json:"terms_size" db:"terms_size"`
	SpikeHeight    int32     `json:"spike_height" db:"spike_height"`
	SpikeType      string    `json:"spike_type" db:"spike_type"`
	AlertOnNewData bool      `json:"alert_on_new_data" db:"alert_on_new_data"`
	ThresholdRef   int32     `json:"threshold_ref" db:"threshold_ref"`
	ThresholdCur   int32     `json:"threshold_cur" db:"threshold_cur"`
}

type WhitelistAlert struct {
	CompareKey string `json:"compare_key" db:"compare_key"`
	Whitelist  string `json:"whitelist" db:"whitelist"`
	IgnoreNull bool   `json:"ignore_null" db:"ignore_null"`
}
