CREATE TABLE IF NOT EXISTS  http_monitor (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	"name" TEXT NOT NULL,
	check_timeout INTEGER NOT NULL,
	fail_trashhold SMALLINT NOT NULL,
	is_enabled BOOLEAN NOT NULL DEFAULT true,
	http_url TEXT NOT NULL,
	http_method TEXT NOT NULL,
	request_headers TEXT NOT NULL,
	valid_response_statues TEXT NOT NULL,
	contact_ids TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
	CONSTRAINT fk_http_monitor_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);