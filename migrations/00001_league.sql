-- +goose Up

CREATE TABLE IF NOT EXISTS `league` (
	`id` int(11) NOT NULL,
	`name` varchar(32) NOT NULL,
	`short_name` varchar(8) DEFAULT NULL,
	`logo_link` varchar(256) DEFAULT NULL,
	`web_site` varchar(128) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `league`
	ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `name` (`name`), ADD UNIQUE KEY `short_name` (`short_name`), ADD UNIQUE KEY `web_site` (`web_site`);

ALTER TABLE `league`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE `league`;
