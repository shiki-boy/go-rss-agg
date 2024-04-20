-- +goose Up
-- +goose StatementBegin
CREATE TABLE Todo (
    id  SERIAL PRIMARY KEY,
    title VARCHAR(255),
    done BOOLEAN,
    created_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Todo;
-- +goose StatementEnd
