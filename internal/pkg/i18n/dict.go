package i18n

import "context"

const defaultLang = "ru"

// Translate ...
func Translate(ctx context.Context, key string, defaultString string) string {
	lang := GetLangFromContext(ctx)

	v, ok := dictionary[lang][key]
	if ok {
		return v
	}
	return defaultString
}

var dictionary = map[string]map[string]string{
	"ru": {
		"certificate_not_found":                       "Сертификат не найден",
		"game_not_found":                              "Игра не найдена",
		"game_result_already_exists_lexeme":           "Результат игры уже есть",
		"game_result_not_found_lexeme":                "Результат игры не найден",
		"invalid_certificate_info_json_value":         "Некорректное значение информации о сертификате",
		"invalid_certificate_type":                    "Некорректный тип сертификата",
		"invalid_date":                                "Некорректная дата",
		"invalid_game_number":                         "Некорректный номер игры",
		"invalid_game_result_result_place_value":      "Некорректное значение результирующего места",
		"invalid_game_result_round_points_json_value": "Некорректное значение очков раундов",
		"invalid_game_type":                           "Некорректный тип игры",
		"invalid_league_id":                           "Некорректная лига",
		"invalid_max_players":                         "Некорректное максимальное число игроков",
		"invalid_place_id":                            "Некорректное место проведения",
		"invalid_price":                               "Некорректная цена",
		"league_not_found":                            "Лига не найдена",
		"lottery_not_available":                       "Лотерея не доступна",
		"lottery_not_implemented":                     "Лотерея не реализована для этой лиги",
		"lottery_permission_denied":                   "Доступ к регистрации в лотерее запрещён",
		"no_free_slots":                               "Нет мест",
		"place_not_found":                             "Место не найдено",
		"spent_on_game_not_found":                     "Игра, на которой потратили сертификат, не найдена",
		"unauthenticated_request":                     "Неавторизованный запрос",
		"user_not_found":                              "Пользователь не найден",
		"user_already_exists":                         "Пользователь уже существует",
		"won_on_game_not_found":                       "Игра, на которой выиграли сертификат, не найдена",

		// errors
		"err_email_validation":         "Некорректный формат email",
		"err_name_alphabet_validation": "Допустим только русский набор букв",
		"err_name_length_validation":   "Длина имени должна быть от 1 до 100 символов",
		"err_name_required_validation": "Имя пользователя обязательно",
		"err_phone_validation":         "Некорректный формат номера телефона",
		"err_state_validation":         "Некорректный стейт пользователя",
	},
}
