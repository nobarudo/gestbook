
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `posts` (
  `id` INT AUTO_INCREMENT NOT NULL,
  `text` VARCHAR(256) NOT NULL,
  PRIMARY KEY(`id`)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `posts`;

