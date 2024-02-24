-- +goose Up

CREATE TABLE IF NOT EXISTS `logs` (
	`id` int(11) NOT NULL,
	`timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`user_id` int(11) DEFAULT NULL,
	`action_id` int(11) NOT NULL,
	`message_id` int(11) NOT NULL,
	`object_type` varchar(32) DEFAULT NULL,
	`object_id` int(11) DEFAULT NULL,
	`metadata` varchar(512) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `logs`
	ADD PRIMARY KEY (`id`); 

ALTER TABLE `logs`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE `logs`;
