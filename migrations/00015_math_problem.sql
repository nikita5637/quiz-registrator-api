-- +goose Up

CREATE TABLE IF NOT EXISTS `math_problem` (
	`id` int(11) NOT NULL,
	`fk_game_id` int(11) NOT NULL,
	`url` varchar(512) NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `math_problem`
	ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `fk_game_id` (`fk_game_id`);

ALTER TABLE `math_problem`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `math_problem`
	ADD CONSTRAINT `math_problem_ibfk_1` FOREIGN KEY (`fk_game_id`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `math_problem`;
