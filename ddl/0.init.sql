CREATE DATABASE IF NOT EXISTS sample;

CREATE TABLE IF NOT EXISTS `product` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(16) COLLATE utf8mb4_bin NOT NULL,
  `price` BIGINT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

TRUNCATE TABLE product;
INSERT INTO product (code, price) values ('A001', 100);
INSERT INTO product (code, price) values ('A002', 140);
INSERT INTO product (code, price) values ('A003', 240);
INSERT INTO product (code, price) values ('A004', 1140);
