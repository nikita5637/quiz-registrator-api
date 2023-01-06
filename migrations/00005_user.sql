-- +goose Up

CREATE TABLE IF NOT EXISTS `user` (
	`id` int(11) NOT NULL,
	`name` varchar(100) NOT NULL,
	`telegram_id` bigint(11) NOT NULL,
	`email` varchar(255) DEFAULT NULL,
	`phone` varchar(12) DEFAULT NULL,
	`state` int(11) NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `user`
	ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `telegram_id` (`telegram_id`), ADD UNIQUE KEY `email` (`email`), ADD UNIQUE KEY `phone` (`phone`);

ALTER TABLE `user`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE `user`;
