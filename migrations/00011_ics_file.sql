-- +goose Up

CREATE TABLE IF NOT EXISTS `ics_file` (
	`id` int(11) NOT NULL,
	`fk_game_id` int(11) NOT NULL,
	`name` VARCHAR(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `ics_file`
	ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `fk_game_id` (`fk_game_id`), ADD UNIQUE KEY `name` (`name`);

ALTER TABLE `ics_file`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `ics_file`
	ADD CONSTRAINT `ics_file_ibfk_1` FOREIGN KEY (`fk_game_id`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `ics_file`;
