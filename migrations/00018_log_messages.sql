-- +goose Up

CREATE TABLE IF NOT EXISTS `log_messages` (
	`id` int(11) NOT NULL,
	`message` varchar(512) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `log_messages`
	ADD PRIMARY KEY (`id`);

ALTER TABLE `log_messages`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

INSERT INTO `log_messages` (`message`) VALUES
	("Выполнена регистрация на игру"),
	("Отменена регистрация на игру"),
	("Изменён тип оплаты за игру"),
	("Получен список всех игр"),
	("Получен список ID игр пользователя"),
	("Получен список прошедших и зарегистрированных игр"),
	("Получен список фотографий с игры");

-- +goose Down

DROP TABLE `log_messages`;
