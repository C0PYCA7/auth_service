-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100),
    surname  VARCHAR(100),
    login    VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    email    VARCHAR(255) UNIQUE
);

CREATE INDEX idx_user_login ON users(login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE users;
-- +goose StatementEnd