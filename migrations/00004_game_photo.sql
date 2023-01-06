-- +goose Up

CREATE TABLE IF NOT EXISTS `game_photo` (
	`id` int(11) NOT NULL,
	`fk_game_id` int(11) NOT NULL,
	`url` varchar(512) NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `game_photo`
	ADD PRIMARY KEY (`id`), ADD KEY `fk_game_id` (`fk_game_id`);

ALTER TABLE `game_photo`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `game_photo`
	ADD CONSTRAINT `game_photo_ibfk_1` FOREIGN KEY (`fk_game_id`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `game_photo`;
