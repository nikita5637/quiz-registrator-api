-- +goose Up

ALTER TABLE `game` ADD `is_in_master` tinyint(1) NOT NULL DEFAULT '0' AFTER `registered`;

-- +goose Down

ALTER TABLE `game` DROP `is_in_master`;