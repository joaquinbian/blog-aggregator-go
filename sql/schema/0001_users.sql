-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id bigserial PRIMARY KEY,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name VARCHAR(50) UNIQUE NOT NULL
)

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd