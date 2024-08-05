-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    first_name varchar(50) NULL,
    last_name varchar(50) NULL,
    username varchar(50) NULL,
    email varchar(255) NULL,
    password varchar(60) NULL,

    confirmation_token varchar(100) NULL,
    confirmation_sent_at timestamptz NULL,
    confirmed_at timestamptz NULL,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);

-- +goose Down
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;
