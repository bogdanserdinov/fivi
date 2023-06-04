-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS comments (
    id      uuid    PRIMARY KEY DEFAULT uuid_generate_v4(),
    text VARCHAR                        NOT NULL,
    post_id uuid    NOT NULL,
    creator_id VARCHAR                     NOT NULL
);
CREATE INDEX comments_creator_id ON comments USING BTREE (creator_id);
-- +migrate Down
DROP INDEX IF EXISTS comments_creator_id;
DROP TABLE IF EXISTS comments;
