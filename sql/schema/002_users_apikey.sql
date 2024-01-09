-- +goose Up
ALTER TABLE users ADD COLUMN api_key varchar(64) unique not null default(
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
ALTER TABLE users DROP C0LUMN api_key;