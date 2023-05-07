CREATE TABLE IF NOT EXISTS articles(
    id BIGSERIAL,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    title VARCHAR(64) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY(id)
);

ALTER TABLE articles ADD CONSTRAINT fk_articles_users FOREIGN KEY(user_id) REFERENCES users (id);

ALTER TABLE articles ADD CONSTRAINT fk_articles_categories FOREIGN KEY(category_id) REFERENCES categories (id);
