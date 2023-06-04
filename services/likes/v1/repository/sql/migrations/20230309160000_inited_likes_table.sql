-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS likes (
    post_id uuid    NOT NULL,
    user_id VARCHAR                     NOT NULL,
    PRIMARY KEY(post_id,user_id )
);
-- +migrate Down
DROP TABLE IF EXISTS likes;
