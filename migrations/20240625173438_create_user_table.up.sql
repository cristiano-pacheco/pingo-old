CREATE TABLE IF NOT EXISTS  users (
	id UUID NOT NULL,
	"name" TEXT NOT NULL,
	email TEXT NOT NULL,
	password_hash BYTEA NOT NULL,
	"status" TEXT NOT NULL,
	reset_password_token BYTEA NULL,
	account_confirmation_token BYTEA NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS user_email_idx ON users (email);