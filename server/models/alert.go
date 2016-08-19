package models

type Alert struct {
	AlertId   int64  `json:"alert_id"`
	EsHost    string `json:"es_host"`
	EsPort    int32  `json:"es_port"`
	EsIndex   string `json:"es_index"`
	RuleName  string `json:"rule_name"`
	RuleType  string `json:"rule_type"`
	AlertType string `json:"alert_type"`
}
