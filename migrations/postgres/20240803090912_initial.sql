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

CREATE TABLE contact_us (
    id serial PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS contact_us;
