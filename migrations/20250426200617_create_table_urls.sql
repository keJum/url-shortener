-- +goose NO TRANSACTION
-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY,
    alias TEXT UNIQUE NOT NULL,
    url TEXT UNIQUE NOT NULL);
Create INDEX IF NOT EXISTS idx_alias ON urls (alias);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX idx_alias;
DROP TABLE urls;
-- +goose StatementEnd
