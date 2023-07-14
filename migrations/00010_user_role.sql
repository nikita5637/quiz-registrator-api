-- +goose Up

CREATE TABLE IF NOT EXISTS `user_role` (
	`id` int(11) NOT NULL,
	`fk_user_id` int(11) NOT NULL,
	`role` ENUM('admin','management','user') NOT NULL DEFAULT 'user',
	`created_at` timestamp NULL DEFAULT NULL,
	`deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `user_role`
	ADD PRIMARY KEY (`id`), ADD KEY `fk_user_id` (`fk_user_id`);

ALTER TABLE `user_role`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `user_role`
	ADD CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`fk_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down

DROP TABLE `user_role`;
