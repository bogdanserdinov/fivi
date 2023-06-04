-- +migrate Up
-- +migrate StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE
OR REPLACE FUNCTION did_mappings_update_updated_at_column() RETURNS TRIGGER AS $$
BEGIN NEW .updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';
-- +migrate StatementEnd
CREATE TABLE IF NOT EXISTS did_mappings (
                                            did           VARCHAR NOT NULL PRIMARY KEY,
                                            long_form_uri VARCHAR NOT NULL,
                                            document      BYTEA   NOT NULL,
                                            updated_at    TIMESTAMP DEFAULT NULL,
                                            created_at    TIMESTAMP NOT NULL DEFAULT now()
    );

CREATE TABLE IF NOT EXISTS did_to_message (
                                              did     VARCHAR PRIMARY KEY NOT NULL,
                                              message VARCHAR NOT NULL
);

CREATE TRIGGER update_did_mappings_modtime BEFORE
    UPDATE ON did_mappings FOR EACH ROW EXECUTE PROCEDURE did_mappings_update_updated_at_column();
-- +migrate Down
DROP TRIGGER IF EXISTS update_did_mappings_modtime ON did_mappings;
DROP TABLE IF EXISTS did_mappings;
DROP FUNCTION IF EXISTS did_mappings_update_updated_at_column();
