-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50),
    url VARCHAR(250),
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd

