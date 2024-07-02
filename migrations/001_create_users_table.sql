-- +migrate Up
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   passport_number VARCHAR(10) NOT NULL UNIQUE,
   name VARCHAR(50) NOT NULL,
   surname VARCHAR(50) NOT NULL,
   patronymic VARCHAR(50),
   address VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS users;