-- +goose Up
CREATE TABLE users(
    user_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_name TEXT unique,
    user_password TEXT
);

-- +goose Down
DROP TABLE users;