ALTER TABLE articles DROP CONSTRAINT IF EXISTS fk_articles_categories;

ALTER TABLE articles DROP CONSTRAINT IF EXISTS fk_articles_users;

DROP TABLE IF EXISTS articles;
