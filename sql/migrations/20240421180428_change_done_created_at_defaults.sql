-- +goose Up
-- +goose StatementBegin
ALTER TABLE Todo ALTER COLUMN done SET DEFAULT false;
ALTER TABLE Todo ALTER COLUMN created_at SET DEFAULT now();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Todo ALTER COLUMN done DROP DEFAULT;
ALTER TABLE Todo ALTER COLUMN created_at DROP DEFAULT;
-- +goose StatementEnd
