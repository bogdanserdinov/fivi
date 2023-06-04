-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS posts (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    payload       VARCHAR NOT NULL DEFAULT '',
    creator_id    uuid   NOT NULL DEFAULT uuid_generate_v4(),
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
-- +migrate Down
DROP TABLE IF EXISTS posts;
