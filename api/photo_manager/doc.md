# Пакет photo_manager

Отвечает за создание и получение фотографий с игр. 

В пакете photo_manager есть **один** сервис: *Service*

Сервис ***Service*** содержит в себе ручки, необходимые для *создания* и *получения* фотографий. Ручки сервиса ***Service*** возвращают ошибки **только** если какие-то ошибки возникают при взаимодействии с базой данных и **не выполняют** никаких проверок бизнес-логики.

---
## Список ручек
[/photo_manager.Service/AddGamePhotos](#/photo_manager.Service/AddGamePhotos)  
[/photo_manager.Service/GetPhotosByGameID](#/photo_manager.Service/GetPhotosByGameID)  

---
## Ручки
В этом разделе перечислены все ручки пакета photo_manager

### <a id="/photo_manager.Service/AddGamePhotos">AddGamePhotos</a>
### Описание
Добавляет список фотографий.
### Путь
`/photo_manager.Service/AddGamePhotos`
### Роли
+ management
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_NOT_FOUND | Игра не найдена |
| Internal | | В остальных случаях |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **management** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/photo_manager.Service/GetPhotosByGameID">GetPhotosByGameID</a>
### Описание
Возвращает список URL ссылок на фотографии с игры.
### Путь
`/photo_manager.Service/GetPhotosByGameID`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| FailedPrecondition | GAME_NOT_FOUND | Игра не найдена |
| Internal | | В остальных случаях |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---