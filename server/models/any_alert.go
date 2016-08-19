package models

type AnyAlert struct {
	AlertId  int64  `json:"alert_id"`
	QueryKey string `json:"query_key"`
}
