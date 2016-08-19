DROP TABLE user;
DROP TABLE alert;
DROP TABLE any_alert;
DROP TABLE blacklist_alert;
DROP TABLE cardinality_alert;
DROP TABLE change_alert;
DROP TABLE flatline_alert;
DROP TABLE frequency_alert;
DROP TABLE new_term_alert;
DROP TABLE spike_alert;
DROP TABLE whitelist_alert;

CREATE TABLE login (
  user_id BIGSERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  salt TEXT NOT NULL,
  email TEXT NOT NULL
);

CREATE TABLE alert (
  alert_id BIGSERIAL PRIMARY KEY,
  es_host TEXT,
  es_port INT,
  es_index TEXT,
  rule_name TEXT,
  rule_type TEXT,
  alert_type TEXT
);

CREATE TABLE any_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT
);

CREATE TABLE blacklist_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  compare_key TEXT,
  blacklist TEXT
);

CREATE TABLE cardinality_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT,
  timeframe TIMESTAMP,
  cardinality_field TEXT,
  max_cardinality BOOL,
  min_cardinality BOOL
);

CREATE TABLE change_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  compare_key TEXT,
  ignore_null BOOL,
  query_key TEXT,
  timeframe TIMESTAMP
);

CREATE TABLE flatline_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT,
  timeframe TIMESTAMP,
  use_count_query BOOL,
  doc_type TEXT,
  threshold INT
);

CREATE TABLE frequency_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT,
  timeframe TIMESTAMP,
  num_events INT,
  attach_related BOOL,
  use_count_query BOOL,
  use_terms_query BOOL,
  doc_type TEXT,
  terms_size INT
);

CREATE TABLE new_term_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT,
  use_terms_query BOOL,
  doc_type TEXT,
  terms_size INT,
  fields TEXT,
  terms_window_size TIMESTAMP,
  window_step_size TIMESTAMP,
  alert_on_missing_fields BOOL
);

CREATE TABLE spike_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  query_key TEXT,
  timeframe TIMESTAMP,
  use_count_query BOOL,
  use_terms_query BOOL,
  doc_type TEXT,
  terms_size INT,
  spike_height INT,
  spike_type TEXT,
  alert_on_new_data BOOL,
  threshold_ref INT,
  threshold_cur INT
);

CREATE TABLE whitelist_alert (
  alert_id BIGINT REFERENCES alert (alert_id),
  compare_key TEXT,
  whitelist TEXT,
  ignore_null BOOL
);





