-- +goose Up

INSERT INTO `log_messages` (`message`) VALUES
	("Получена информация об игре"),
	("Получен список сертификатов"),
	("Получен признак того, есть ли фотографии у игры");

-- +goose Down

DELETE FROM `log_messages` WHERE `log_messages`.`message` = "Получена информация об игре";
DELETE FROM `log_messages` WHERE `log_messages`.`message` = "Получен список сертификатов";
DELETE FROM `log_messages` WHERE `log_messages`.`message` = "Получен признак того, есть ли фотографии у игры";

ALTER TABLE `log_messages` AUTO_INCREMENT=8;
