
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int not null, name varchar(30) not null);
-- +migrate Down
DROP TABLE IF EXISTS users;