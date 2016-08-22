INSERT INTO login (username, password, salt, email) VALUES ('admin', 'admin', 'abc123', 'a@b.c');

INSERT INTO alert (es_host, es_port, es_index, rule_name, rule_type, alert_type, alert_options) VALUES ('localhost', '9200', 'topbeat-*', 'CPU', 'frequency', 'email', '{"email": "a@b.c"}');
