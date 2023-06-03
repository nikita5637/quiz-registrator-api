-- +goose Up

ALTER TABLE `game_result` ADD `points` VARCHAR(256) NULL DEFAULT NULL ;

-- +goose Down

ALTER TABLE `game_result` DROP `points`;
