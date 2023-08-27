# Сервис динамического сегментирования пользователей

### Запуск

Для запуска требуется добавить файл .env, заполнить его в соответствии с:
`DB_USER= ...`
`DB_PASS= ...`
`DB_NAME= ...`


### Использование
Запустить сервер можно командой `docker compose up`


### Примеры запросов и ответов
### Интервал времени между каждым запросом менее одной минуты

#### Сегменты
##### Добавление нового сегмента
##### POST запрос на адрес `localhost:80/add`
Запрос
```json
{
    "segment":"AVITO_DISCOUNT_70"
}
```
Ответ
```json
{
    "segment":"AVITO_DISCOUNT_70"
}
```
&nbsp;
Запрос (добавление уже существующего элемента)
```json
{
    "segment":"AVITO_DISCOUNT_70"
}
```
Ответ
```json
{
    "message": "Segment you are trying to add already exists."
}
```
---
##### Вывод всех доступных сегментов
##### GET запрос на адрес `localhost:80/show`
Ответ (для наглядности добавил все сегменты)
```json
[
    {
        "segment": "AVITO_DISCOUNT_50"
    },
    {
        "segment": "AVITO_VOICE_MESSAGES"
    },
    {
        "segment": "AVITO_PERFORMANCE_VAS"
    },
    {
        "segment": "AVITO_DISCOUNT_70"
    }
]
```
---
##### Удаление сегмента
##### DELETE запрос на адрес `localhost:80/del`
Запрос
```json
{
    "segment":"AVITO_DISCOUNT_70"
}
```
Ответ
```json
{
    "message": "Segment deleted successfully."
}
```
&nbsp;
Запрос (несуществующего сегмента)
```json
{
    "segment":"sample_segment"
}
```
Ответ
```json
{
    "message": "there is no segment like sample_segment"
}
```
---
#### Пользователи
##### Добавление / изменение пользователя
##### POST запрос на адрес `localhost:80/user`
Запрос на добавление пользователя без сегментов
```json
{
    "id": 1000
}
```
Ответ
```json
{
    "User": {
        "id": 1000,
        "segments": null
    },
    "message": "New user created."
}
```
&nbsp;
Запрос
```json
{
    "id": 1001,
    "addseg": ["AVITO_DISCOUNT_70", "AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"]
}
```
Ответ
```json
{
    "User": {
        "id": 1001,
        "segments": [
            "AVITO_DISCOUNT_70",
            "AVITO_VOICE_MESSAGES",
            "AVITO_PERFORMANCE_VAS"
        ]
    },
    "message": "New user created."
}
```
&nbsp;
Запрос (добавление пользователя с наличием сегментов, которых нет в списке добавленных)
```json
{
    "id": 1002,
    "addseg": ["AVITO_DISCOUNT_70", "AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "this wont be added"]
}
```
Ответ
```json
{
    "User": {
        "id": 1002,
        "segments": [
            "AVITO_DISCOUNT_70",
            "AVITO_VOICE_MESSAGES",
            "AVITO_PERFORMANCE_VAS"
        ]
    },
    "message": "New user created. One or more segments are not in the list."
}
```
&nbsp;
Запрос на удаление сегментов пользователя
```json
{
    "id": 1002,
    "delseg": ["AVITO_DISCOUNT_70"]
}
```
Ответ
```json
{
    "User": {
        "id": 1002,
        "segments": [
            "AVITO_VOICE_MESSAGES",
            "AVITO_DISCOUNT_50",
            "AVITO_PERFORMANCE_VAS"
        ]
    },
    "message": "Users segments are changed."
}
```
&nbsp;
Запрос на одновременное добавление и удаление сегментов пользователя
```json
{
    "id": 1002,
    "addseg": ["AVITO_DISCOUNT_50"],
    "delseg": ["AVITO_PERFORMANCE_VAS"]
}
```
Ответ
```json
{
    "User": {
        "id": 1002,
        "segments": [
            "AVITO_DISCOUNT_70",
            "AVITO_VOICE_MESSAGES",
            "AVITO_DISCOUNT_50"
        ]
    },
    "message": "Users segments are changed."
}
```
&nbsp;
Запрос на добавление сегментов, которые уже есть у пользователя / которые уже указаны в запросе
```json
{
    "id": 1002,
    "addseg": ["AVITO_DISCOUNT_50", "AVITO_PERFORMANCE_VAS", "AVITO_PERFORMANCE_VAS"]
}
```
Ответ
```json
{
    "User": {
        "id": 1002,
        "segments": [
            "AVITO_DISCOUNT_70",
            "AVITO_VOICE_MESSAGES",
            "AVITO_DISCOUNT_50",
            "AVITO_PERFORMANCE_VAS"
        ]
    },
    "message": "Users segments are changed."
}
```
---
##### Вывод всех пользователей
##### GET запрос на адрес `localhost:80/user/show`
Ответ
```json
[
    {
        "id": 1000,
        "segments": null
    },
    {
        "id": 1001,
        "segments": [
            "AVITO_DISCOUNT_70",
            "AVITO_VOICE_MESSAGES",
            "AVITO_PERFORMANCE_VAS"
        ]
    },
    {
        "id": 1002,
        "segments": [
            "AVITO_VOICE_MESSAGES",
            "AVITO_DISCOUNT_50",
            "AVITO_PERFORMANCE_VAS"
        ]
    }
]
```
---
##### Вывод выбранного пользователя
##### GET запрос на адрес `localhost:80/user/show/1001`
Ответ
```json
{
    "id": 1001,
    "segments": [
        "AVITO_DISCOUNT_70",
        "AVITO_VOICE_MESSAGES",
        "AVITO_PERFORMANCE_VAS"
    ]
}
```
---
##### Удаление выбранного пользователя
##### DELETE запрос на адрес `localhost:80/user/delete/1001`
Ответ
```json
{
    "message": "User deleted successfully."
}
```
&nbsp;
##### DELETE запрос на адрес с указанием ID несуществующего пользователя `localhost:80/user/delete/1111`
Ответ
```json
{
    "message": "there is no ID like 1111"
}
```