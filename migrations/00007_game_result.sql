-- +goose Up

CREATE TABLE IF NOT EXISTS `game_result` (
	`id` int(11) NOT NULL,
	`fk_game_id` int(11) NOT NULL,
	`place` tinyint(3) unsigned NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `game_result`
	ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `game_result_ibfk_1` (`fk_game_id`);

ALTER TABLE `game_result`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `game_result`
	ADD CONSTRAINT `game_result_ibfk_1` FOREIGN KEY (`fk_game_id`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `game_result`;
