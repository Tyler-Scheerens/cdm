package models

import (
  "github.com/jmoiron/sqlx/types"
)

type Alert struct {
	AlertId      int64  `json:"alert_id" db:"alert_id"`
	EsHost       string `json:"es_host" db:"es_host"`
	EsPort       int32  `json:"es_port" db:"es_port"`
	EsIndex      string `json:"es_index" db:"es_index"`
	RuleName     string `json:"rule_name" db:"rule_name"`
	RuleType     string `json:"rule_type" db:"rule_type"`
	AlertType    string `json:"alert_type" db:"alert_type"`
  AlertOptions types.JSONText `json:"alert_options" db:"alert_options"`
}
