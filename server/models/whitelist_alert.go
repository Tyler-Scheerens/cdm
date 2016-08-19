package models

type WhitelistAlert struct {
	AlertId    int64  `json:"alert_id"`
	CompareKey string `json:"compare_key"`
	Whitelist  string `json:"whitelist"`
	IgnoreNull bool   `json:"ignore_null"`
}
