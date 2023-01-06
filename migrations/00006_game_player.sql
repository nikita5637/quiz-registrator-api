-- +goose Up

CREATE TABLE IF NOT EXISTS `game_player` (
	`id` int(11) NOT NULL,
	`fk_game_id` int(11) NOT NULL,
	`fk_user_id` int(11) DEFAULT NULL,
	`registered_by` int(11) NOT NULL,
	`degree` tinyint(3) unsigned NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `game_player`
	ADD PRIMARY KEY (`id`), ADD KEY `fk_game_id` (`fk_game_id`), ADD KEY `fk_user_id` (`fk_user_id`);

ALTER TABLE `game_player`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `game_player`
	ADD CONSTRAINT `game_player_ibfk_1` FOREIGN KEY (`fk_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	ADD CONSTRAINT `game_player_ibfk_2` FOREIGN KEY (`fk_game_id`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `game_player`;
