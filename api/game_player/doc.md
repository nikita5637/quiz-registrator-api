# Пакет game_player

Отвечает за создание, удаление, изменение и получение игроков. Игрок и пользователь это разные сущности. Один пользователь может быть игроком на нескольких играх.

В пакете game_player есть **два** сервиса: *Service* и *RegistratorService*.

Сервис ***Service*** содержит в себе ручки, необходимые для *создания*, *удаления*, *изменения* и *получения* игроков. Ручки сервиса ***Service*** возвращают ошибки **только** если какие-то ошибки возникают при взаимодействии с  базой данных и **не выполняют** никаких проверок бизнес-логики. Например ручка **создания игрока** не выполняет проверку того, что игра уже прошла.

Сервис ***RegistratorService*** наоборот **выполняет** дополнительные проверки бизнес-логики и не выполняет запрос, если произошли какие-то ошибки бизнес-логики. Например ручка **регистрации игрока** не зарегистрирует игрока на игру, которая уже прошла.

---
## Константы
```
DEGREE_INVALID = 0 - некорректное значение
DEGREE_LIKELY = 1 - игрок точно придёт на игру
DEGREE_UNLIKELY = 2 - игрок может быть придёт на игру
```

---
## Список ручек
[/game_player.Service/CreateGamePlayer](#/game_player.Service/CreateGamePlayer)  
[/game_player.Service/DeleteGamePlayer](#/game_player.Service/DeleteGamePlayer)  
[/game_player.Service/GetGamePlayer](#/game_player.Service/GetGamePlayer)  
[/game_player.Service/GetGamePlayersByGameID](#/game_player.Service/GetGamePlayersByGameID)  
[/game_player.Service/PatchGamePlayer](#/game_player.Service/PatchGamePlayer)  
[/game_player.RegistratorService/RegisterPlayer](#/game_player.RegistratorService/RegisterPlayer)  
[/game_player.RegistratorService/UnregisterPlayer](#/game_player.RegistratorService/UnregisterPlayer)

---
## Ручки
В этом разделе перечислены все ручки пакета game_player

### <a id="/game_player.Service/CreateGamePlayer">CreateGamePlayer</a>
### Описание
Создаёт сущность игрока.
### Путь
`/game_player.Service/CreateGamePlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | GAME_PLAYER_ALREADY_EXISTS | Игрок уже существует |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | GAME_NOT_FOUND | Игра не найдена |
| InvalidArgument | INVALID_DEGREE | Некорректное значение вероятности |
| InvalidArgument | INVALID_GAME_ID | Некорректное значение ID игры |
| InvalidArgument | INVALID_REGISTERED_BY | Некорректное значение ID пользователя, регистрирующего игрока |
| InvalidArgument | INVALID_USER_ID | Некорректное значение ID пользователя |
| InvalidArgument | USER_NOT_FOUND | Пользователь не найден|
| - | - | - |

---
### <a id="/game_player.Service/DeleteGamePlayer">DeleteGamePlayer</a>
### Описание
Удаляет сущность игрока по её ID.
### Путь
`/game_player.Service/DeleteGamePlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| NotFound | GAME_PLAYER_NOT_FOUND | Игрок не найден |
| - | - | - |

---
### <a id="/game_player.Service/GetGamePlayer">GetGamePlayer</a>
### Описание
Возвращает сущность игрока по её ID.
### Путь
`/game_player.Service/GetGamePlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| NotFound | GAME_PLAYER_NOT_FOUND | Игрок не найден |
| - | - | - |

---
### <a id="/game_player.Service/GetGamePlayersByGameID">GetGamePlayersByGameID</a>
### Описание
Возвращает список сущностей игроков по ID игры.
### Путь
`/game_player.Service/GetGamePlayersByGameID`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В любом случае |
| - | - | - |

---
### <a id="/game_player.Service/PatchGamePlayer">/game_player.Service/PatchGamePlayer</a>
### Описание
Изменяет сущность игрока по ID сущности.
### Путь
`/game_player.Service/PatchGamePlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | GAME_PLAYER_ALREADY_EXISTS | Игрок уже существует |
| Internal |  | В остальных случаях |
| NotFound | GAME_PLAYER_NOT_FOUND | Игрок не найден |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | GAME_NOT_FOUND | Игра не найдена |
| InvalidArgument | INVALID_DEGREE | Некорректное значение вероятности |
| InvalidArgument | INVALID_GAME_ID | Некорректное значение ID игры |
| InvalidArgument | INVALID_REGISTERED_BY | Некорректное значение ID пользователя, регистрирующего игрока |
| InvalidArgument | INVALID_USER_ID | Некорректное значение ID пользователя |
| InvalidArgument | USER_NOT_FOUND | Пользователь не найден |
| - | - | - |

---
### <a id="/game_player.RegistratorService/RegisterPlayer">RegisterPlayer</a>
### Описание
Регистрирует пользователя на игру. Выполняет внутренние проверки бизнес-логики перед регистрацией.
### Путь
`/game_player.RegistratorService/RegisterPlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | GAME_PLAYER_ALREADY_REGISTERED | Игрок уже зарегистрирован |
| FailedPrecondition | GAME_HAS_PASSED | Игра прошла |
| FailedPrecondition | GAME_NOT_FOUND | Игра не найдена |
| FailedPrecondition | THERE_ARE_NO_FREE_SLOT | Нет совобдных слотов |
| FailedPrecondition | USER_NOT_FOUND | Пользователь не найден |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | INVALID_DEGREE | Некорректное значение вероятности |
| InvalidArgument | INVALID_GAME_ID | Некорректное значение ID игры |
| InvalidArgument | INVALID_REGISTERED_BY | Некорректное значение ID пользователя, регистрирующего игрока |
| InvalidArgument | INVALID_USER_ID | Некорректное значение ID пользователя |
| - | - | - |

---
### <a id="/game_player.RegistratorService/UnregisterPlayer">UnregisterPlayer</a>
### Описание
Отменяет регистрацию игрока. Выполняет внутренние проверки бизнес-логики перед отменой регистрации.
### Путь
`/game_player.RegistratorService/UnregisterPlayer`
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_HAS_PASSED | Игра прошла |
| FailedPrecondition | GAME_NOT_FOUND | Игра не найдена |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | INVALID_DEGREE | Некорректное значение вероятности |
| InvalidArgument | INVALID_GAME_ID | Некорректное значение ID игры |
| InvalidArgument | INVALID_REGISTERED_BY | Некорректное значение ID пользователя, регистрирующего игрока |
| InvalidArgument | INVALID_USER_ID | Некорректное значение ID пользователя |
| NotFound | | Список игроков пользователя пуст |
| NotFound | GAME_PLAYER_NOT_FOUND | Игрок не найден |
| - | - | - |

---