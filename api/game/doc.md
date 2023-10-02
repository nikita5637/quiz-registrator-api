# Пакет game

Отвечает за создание, удаление, изменение и получение игр.

Проверка на существование игры выполняется путем поиска связки аттрибутов "внешний ID" + "лига" + "место" + "номер" + "дата и время". Если существует запись в базе данных с такими же значениями данных аттрибутов, то считается, что такая игра существует в базе данных.

В пакете game есть **два** сервиса: *Service* и *RegistratorService*.

Сервис ***Service*** содержит в себе ручки, необходимые для *создания*, *удаления*, *изменения* и *получения* игр. Ручки сервиса ***Service*** возвращают ошибки **только** если какие-то ошибки возникают при взаимодействии с  базой данных и **не выполняют** никаких проверок бизнес-логики.

Сервис ***RegistratorService*** **выполняет** дополнительные проверки бизнес-логики и не выполняет запрос, если произошли ошибки бизнес-логики. Например ручка **регистрации игры** не выполнит регистрацию, если игра уже прошла. 

---
## Константы
```
GAME_TYPE_INVALID = 0 - некорректное значение
GAME_TYPE_CLASSIC = 1 - классическая игра
GAME_TYPE_THEMATIC = 2 - тематическая игра
GAME_TYPE_MOVIES_AND_MUSIC = 5 - КиМ игра
GAME_TYPE_CLOSED = 6 - закрытая игра
GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC = 9 - тематическая КиМ игра

PAYMENT_INVALID = 0 - некорректное значение
PAYMENT_CASH = 1 - оплата деньгами
PAYMENT_CERTIFICATE = 2 - оплата проходкой
PAYMENT_MIXED = 3 - оплата и деньгами и проходкой
```

---
## Список ручек
[/game.Service/BatchGetGames](#/game.Service/BatchGetGames)  
[/game.Service/CreateGame](#/game.Service/CreateGame)  
[/game.Service/DeleteGame](#/game.Service/DeleteGame)  
[/game.Service/GetGame](#/game.Service/GetGame)  
[/game.Service/ListGames](#/game.Service/ListGames)  
[/game.Service/PatchGame](#/game.Service/PatchGame)  
[/game.Service/SearchGamesByLeagueID](#/game.Service/SearchGamesByLeagueID)  
[/game.Service/SearchPassedAndRegisteredGames](#/game.Service/SearchPassedAndRegisteredGames)  
[/game.RegistratorService/RegisterGame](#/game.RegistratorService/RegisterGame)  
[/game.RegistratorService/UnregisterGame](#/game.RegistratorService/UnregisterGame)  
[/game.RegistratorService/UpdatePayment](#/game.RegistratorService/UpdatePayment)  

---
## Ручки
В этом разделе перечислены все ручки пакета game

### <a id="/game.Service/BatchGetGames">BatchGetGames</a>
### Описание
Возвращает список игр по их ID. Сортировка результата **не предусмотрена** и **не гарантируется**.
### Путь
`/game.Service/BatchGetGames`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/CreateGame">CreateGame</a>
### Описание
Создаёт новую игру.
### Путь
`/game.Service/CreateGame`
### Роли
+ management
+ s2s
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | GAME_ALREADY_EXISTS | Игра уже существует |
| FailedPrecondition | LEAGUE_NOT_FOUND | Лига не найдена |
| FailedPrecondition | PLACE_NOT_FOUND | Место проведения игры не найдено |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | INVALID_EXTERNAL_ID | Некорректный ID игры |
| InvalidArgument | INVALID_GAME_DATE | Некорректная дата игры |
| InvalidArgument | INVALID_GAME_NAME | Некорректное название игры |
| InvalidArgument | INVALID_GAME_NUMBER | Некорректный номер игры |
| InvalidArgument | INVALID_GAME_TYPE | Некорректный тип игры |
| InvalidArgument | INVALID_LEAGUE_ID | Некорректный ID лиги |
| InvalidArgument | INVALID_MAX_PLAYERS | Некорректное количество игроков |
| InvalidArgument | INVALID_PAYMENT | Некорректный тип оплаты |
| InvalidArgument | INVALID_PAYMENT_TYPE | Некорректный способ оплаты |
| InvalidArgument | INVALID_PLACE_ID | Некорректный ID места игры |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролями **management** или **s2s** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/DeleteGame">DeleteGame</a>
### Описание
Удаляет игру по её ID.
### Путь
`/game.Service/DeleteGame`
### Роли
+ management
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **management** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/GetGame">GetGame</a>
### Описание
Возвращает игру по её ID.
### Путь
`/game.Service/GetGame`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/ListGames">ListGames</a>
### Описание
Возвращает все игры из базы.  
Гарантируется сортировка по дате игры по возрастанию.
### Путь
`/game.Service/ListGames`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/PatchGame">PatchGame</a>
### Описание
Изменяет сущность игры по её ID.
### Путь
`/game.Service/PatchGame`
### Роли
+ management
+ s2s
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | GAME_ALREADY_EXISTS | Игра уже существует |
| FailedPrecondition | LEAGUE_NOT_FOUND | Лига не найдена |
| FailedPrecondition | PLACE_NOT_FOUND | Место проведения игры не найдено |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | INVALID_EXTERNAL_ID | Некорректный ID игры |
| InvalidArgument | INVALID_GAME_DATE | Некорректная дата игры |
| InvalidArgument | INVALID_GAME_NAME | Некорректное название игры |
| InvalidArgument | INVALID_GAME_NUMBER | Некорректный номер игры |
| InvalidArgument | INVALID_GAME_TYPE | Некорректный тип игры |
| InvalidArgument | INVALID_LEAGUE_ID | Некорректный ID лиги |
| InvalidArgument | INVALID_MAX_PLAYERS | Некорректное количество игроков |
| InvalidArgument | INVALID_PAYMENT | Некорректный тип оплаты |
| InvalidArgument | INVALID_PAYMENT_TYPE | Некорректный способ оплаты |
| InvalidArgument | INVALID_PLACE_ID | Некорректный ID места игры |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролями **management** или **s2s** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/SearchGamesByLeagueID">SearchGamesByLeagueID</a>
### Описание
Возвращает все игры лиги.  
Гарантируется сортировка по дате игры по возрастанию.
### Путь
`/game.Service/SearchGamesByLeagueID`
### Роли
+ public
+ s2s
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.Service/SearchPassedAndRegisteredGames">SearchPassedAndRegisteredGames</a>
### Описание
Возвращает прошедшие зарегистрированные игры(отыгранные игры).  
Параметр **page** указывает на номер страницы и принимает значения [1:)  
Параметр **page_size** указывает на количество игр на странице и принимает значения [1:)  
Гарантируется сортировка по дате игры по убыванию.
### Путь
`/game.Service/SearchGamesByLeagueID`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.RegistratorService/RegisterGame">RegisterGame</a>
### Описание
Регистрирует игру по её ID.
### Путь
`/game.RegistratorService/RegisterGame`
### Роли
+ user
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_HAS_PASSED | Игра уже прошла |
| Internal | | В остальных случаях |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **user** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/game.RegistratorService/UnregisterGame">UnregisterGame</a>
### Описание
Отменяет регистрацию игры по её ID.
### Путь
`/game.RegistratorService/UnregisterGame`
### Роли
+ user
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_HAS_PASSED | Игра уже прошла |
| Internal | | В остальных случаях |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **user** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---

### <a id="/game.RegistratorService/UpdatePayment">UpdatePayment</a>
### Описание
Обновляет тип оплаты за игру по её ID.
### Путь
`/game.RegistratorService/UpdatePayment`
### Роли
+ user
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_HAS_PASSED | Игра уже прошла |
| Internal | | В остальных случаях |
| InvalidArgument | INVALID_PAYMENT | Некорректный тип оплаты |
| NotFound | GAME_NOT_FOUND | Игра не найдена |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **user** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---