# Пакет math_problem

Отвечает за создание и получение математический задач. 

В пакете math_problem есть **один** сервис: *Service*

Сервис ***Service*** содержит в себе ручки, необходимые для *создания* и *получения* математических задач. Ручки сервиса ***Service*** возвращают ошибки **только** если какие-то ошибки возникают при взаимодействии с базой данных и **не выполняют** никаких проверок бизнес-логики.

---
## Список ручек
[/math_problem.Service/CreateMathProblem](#/math_problem.Service/CreateMathProblem)  
[/math_problem.Service/SearchMathProblemByGameID](#/math_problem.Service/SearchMathProblemByGameID)  

---
## Ручки
В этом разделе перечислены все ручки пакета math_problem

### <a id="/math_problem.Service/CreateMathProblem">CreateMathProblem</a>
### Описание
Добавляет математическую задачу.
### Путь
`/math_problem.Service/CreateMathProblem`
### Роли
+ management
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| AlreadyExists | MATH_PROBLEM_ALREADY_EXISTS | Математическая задача уже существует |
| FailedPrecondition | GAME_NOT_FOUND | Игра не найдена |
| Internal | | В остальных случаях |
| InvalidArgument | | Пришёл пустой запрос или произошла ошибка валидатора |
| InvalidArgument | INVALID_GAME_ID | Некорректный ID игры |
| InvalidArgument | INVALID_URL | Некорректный URL |
| PermissionDenied | PERMISSION_DENIED | Вызывающий не владеет ролью **management** |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---
### <a id="/math_problem.Service/SearchMathProblemByGameID">SearchMathProblemByGameID</a>
### Описание
Возвращает URL ссылку на фотографию математической задачи с игры.
### Путь
`/math_problem.Service/SearchMathProblemByGameID`
### Роли
+ public
### Возвращаемые ошибки
| Код | Причина | Описание |
| - | - | - |
| Internal | | В остальных случаях |
| NotFound | MATH_PROBLEM_NOT_FOUND | Математическая задача не найдена |
| PermissionDenied | YOU_ARE_BANNED | Вызывающий забанен |
| - | - | - |

---