CREATE TABLE IF NOT EXISTS users(
    id SERIAL,
    username VARCHAR(16) NOT NULL,
    name VARCHAR(48) NOT NULL,
    password VARCHAR(96) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY(id)
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_users_username ON users(LOWER(username));
