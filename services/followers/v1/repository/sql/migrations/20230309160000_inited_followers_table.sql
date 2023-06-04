-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS followers (
    id uuid PRIMARY KEY NOT NULL,
    "follower_id" VARCHAR NOT NULL,
    "followee_id" VARCHAR NOT NULL
);
CREATE INDEX followers_follower_id ON followers USING BTREE (follower_id);
CREATE INDEX followers_followee_id ON followers USING BTREE (followee_id);

-- +migrate Down
DROP TABLE IF EXISTS followers;
DROP INDEX IF EXISTS followers_follower_id;
DROP INDEX IF EXISTS followers_followee_id;
