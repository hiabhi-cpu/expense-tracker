-- +goose Up
CREATE TABLE money(
    mon_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    mon_desc TEXT ,
    amt INT,
    user_id INT,
    mon_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE money;