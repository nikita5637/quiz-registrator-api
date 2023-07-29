-- +goose Up

ALTER TABLE `game_player` ADD `updated_at` timestamp NULL DEFAULT NULL AFTER `created_at`;

-- +goose Down

ALTER TABLE `game_player` DROP `updated_at`;