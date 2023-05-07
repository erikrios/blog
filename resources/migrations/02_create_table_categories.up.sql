CREATE TABLE IF NOT EXISTS categories(
    id SERIAL,
    name VARCHAR(24) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY(id)
);

CREATE INDEX IF NOT EXISTS idx_categories_name ON categories(name);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_categories_name ON categories(LOWER(name));
