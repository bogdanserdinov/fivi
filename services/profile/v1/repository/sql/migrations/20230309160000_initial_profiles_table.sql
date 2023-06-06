-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS profiles (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR NOT NULL DEFAULT 'fivi_user',
    username      VARCHAR   NOT NULL DEFAULT 'fivi_user',
    email         VARCHAR   NOT NULL DEFAULT 'fivi_user',
    mnemonic      VARCHAR NOT NULL DEFAULT '',
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
-- +migrate Down
DROP TABLE IF EXISTS profiles;
