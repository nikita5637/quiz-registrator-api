-- +goose Up

CREATE TABLE IF NOT EXISTS `certificate` (
	`id` int(11) NOT NULL,
	`type` tinyint(3) unsigned NOT NULL DEFAULT 0,
	`won_on` int(11) NOT NULL,
	`spent_on` int(11) DEFAULT NULL,
	`info` varchar(256) DEFAULT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`updated_at` timestamp NULL DEFAULT NULL,
	`deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `certificate`
	ADD PRIMARY KEY (`id`), ADD KEY `won_on` (`won_on`), ADD KEY `spent_on` (`spent_on`);

ALTER TABLE `certificate`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `certificate`
	ADD CONSTRAINT `game_id_fk_1` FOREIGN KEY (`won_on`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	ADD CONSTRAINT `game_id_fk_2` FOREIGN KEY (`spent_on`) REFERENCES `game` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `certificate`;