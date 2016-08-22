package models

type BlacklistAlert struct {
	AlertId    int64  `json:"alert_id" db:"alert_id"`
	CompareKey string `json:"compare_key" db:"compare_key"`
	Blacklist  string `json:"blacklist" db:"blacklist"`
}
