package models

type Elasticsearch struct {
	ElasticsearchId int64  `json:"elasticsearch_id" db:"elasticsearch_id"`
	IP              string `json:"ip" db:"ip"`
	Port            int32 `json:"port" db:"port"`
	AlertIndex      string `json:"alert_index" db:"alert_index"`
}
