package models

type WhitelistAlert struct {
	AlertId    int64  `json:"alert_id" db:"alert_id"`
	CompareKey string `json:"compare_key" db:"compare_key"`
	Whitelist  string `json:"whitelist" db:"whitelist"`
	IgnoreNull bool   `json:"ignore_null" db:"ignore_null"`
}
