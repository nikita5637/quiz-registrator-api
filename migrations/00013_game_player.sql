-- +goose Up

ALTER TABLE `game_player` ADD INDEX(`registered_by`);
ALTER TABLE `game_player` ADD CONSTRAINT `game_player_ibfk_3` FOREIGN KEY (`registered_by`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

ALTER TABLE `game_player` DROP FOREIGN KEY `game_player_ibfk_3`;
ALTER TABLE `game_player` DROP INDEX `registered_by`;