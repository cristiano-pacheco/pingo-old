CREATE TABLE IF NOT EXISTS  contact (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	"name" TEXT NOT NULL,
	contact_type TEXT NOT NULL,
	contact_data TEXT NOT NULL,
	is_enabled BOOLEAN NOT NULL default TRUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
	CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS user_contact_data_idx ON contact (user_id, contact_data);