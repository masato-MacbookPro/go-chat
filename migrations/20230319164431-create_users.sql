-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    `id` BIGINT NOT NULL AUTO_INCREMENT, 
    `name` VARCHAR(30) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS users;