-- +goose Up

ALTER TABLE `user` ADD `birthdate` DATE NULL DEFAULT NULL AFTER `state`;
ALTER TABLE `user` ADD `sex` TINYINT(3) NULL DEFAULT NULL COMMENT '1 - male, 2 - female' AFTER `birthdate`;

-- +goose Down

ALTER TABLE `user` DROP `birthdate`;
ALTER TABLE `user` DROP `sex`;
