-- +goose Up

CREATE TABLE IF NOT EXISTS `game` (
	`id` int(11) NOT NULL,
	`external_id` int(11) DEFAULT NULL,
	`league_id` int(11) NOT NULL,
	`type` tinyint(3) unsigned NOT NULL,
	`number` varchar(32) NOT NULL,
	`name` varchar(64) DEFAULT NULL,
	`place_id` int(11) NOT NULL,
	`date` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
	`price` int(11) unsigned NOT NULL,
	`payment_type` set('cash','card') DEFAULT NULL,
	`max_players` tinyint(3) unsigned NOT NULL,
	`payment` tinyint(3) unsigned DEFAULT NULL,
	`registered` tinyint(1) NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`updated_at` timestamp NULL DEFAULT NULL,
	`deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `game`
	ADD PRIMARY KEY (`id`);

ALTER TABLE `game`
	ADD KEY `league` (`league_id`);

ALTER TABLE `game`
	ADD KEY `place_id` (`place_id`);

ALTER TABLE `game`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `game`
	ADD CONSTRAINT `place_ibfk_1` FOREIGN KEY (`place_id`) REFERENCES `place` (`id`),
	ADD CONSTRAINT `league_ibfk_1` FOREIGN KEY (`league_id`) REFERENCES `league` (`id`);

-- +goose Down

DROP TABLE `game`;