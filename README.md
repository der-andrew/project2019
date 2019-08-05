# Thesaurus

## Преимущества управление инфраструктурой и конфигурацией:
- Весь код хранится в системе контроля версий (СКВ).
- Единый источник изменений - это СКВ.
- Использование практик разработки программного обеспечения (ПО).
- Повторяемость кода и конфигураций.


## Назначение:
- Подсистема хранения и управления нормативно-справочной информацией

## Возможности:
- CRUD HTTP Restful API
- Механизм обновления инстансов
- Импорт из CSV
- Статические справочники
- Пользовательские справочники
- Гибридные справочники

## Архитектура:

```
                                ____
              -----------      (    )
user 8-) ---> | MainAPP | ---> ( DB )
              -----------      (    )
                  ^             ~~~~
                  |
              ____|_____
              |  CSV   |
              | source |
              ~~~~~~~~~~
```
В качестве DB используется MongoDB.


## API Reference
- GET /status Статус инстанса
- GET /documents Коллекция документов
- POST /documents Создание документа
- PUT /documents Редактирование документа
- DELETE /documents Удаление документа

## Логирование
```
| Контекст | Уровень | Сообщение                                            | Описание                                               |
|----------|---------|------------------------------------------------------|--------------------------------------------------------|
| CORE     | INFO    | resource=DB addr=[MONGO_HOST:PORT] status=CONNECTED  | Соединение с БД успешно установлено                    |
| CORE     | FATAL   | resource=DB addr=[MONGO_HOST:PORT] status=FAILED     | Соединение с БД  не может быть установлено             |
| CORE     | DEBUG   |                                                      | Детализация ошибки подключения к БД                    |
| CORE     | INFO    | version=v[APP_VERSION] status=STARTED                | Приложение успешно запущено                            |
| DOCUMENT | INFO    | documentType=[DOC_TYPE] updateStatus=[UPDATE_STATUS] | Статус обновления справочника (UPDATING|UPDATED|ERROR) |
```

## Параметры запуска
```
-c [путь к файлу конфигурации]
--update - Обновление справочников
```
## Конфигурация
\* - обязательные параметры
_- значения по умолчанию
```

| Параметр             | Тип     | Описание                                                     | ENV                            |
|----------------------|---------|--------------------------------------------------------------|--------------------------------|
| *http.host           | string  | Адрес   бинд-хоста HTTP интерфейса (_0.0.0.0)                |  THESAURUS_HTTP_HOST           |
| http.port            | integer | Порт   бинд-хоста HTTP интерфейса (_80)                      |  THESAURUS_HTTP_PORT           |
| *db.host             | integer | Адрес   хоста БД                                             |  THESAURUS_DB_HOST             |
| db.port              | string  | Порт   хоста БД (_27017)                                     |  THESAURUS_DB_PORT             |
| *db.database         | string  | Название   БД                                                |  THESAURUS_DB_DATABASE         |
| db.login             | string  | Логин   пользователя БД                                      |  THESAURUS_DB_LOGIN            |
| db.password          | string  | Пароль   пользователя БД                                     |  THESAURUS_DB_PASSWORD         |
| csv.path             | string  | Путь до   папки со словарями                                 |  THESAURUS_CSV_PATH            |
| csv.separator.column | string  | Разделитель   полей CSV (,)                                  | THESAURUS_CSV_SEPARATOR_COLUMN |
| sentryDSN            | string  | DSN   аггрегатора ошибок                                     |  THESAURUS_SENTRY_DSN          |
| logging.output       | string  | Вывод   ошибок в _STDOUT или файл (указывается путь до файла)|  THESAURUS_LOGGING_OUTPUT      |
| logging.level        | string  | Уровень   логирования (DEBUG|ERROR|FATAL|_INFO)              |  THESAURUS_LOGGING_LEVEL       |
| logging.format       | string  | Формат   логов (_TEXT|JSON)                                  |  THESAURUS_LOGGING_FORMAT      |
```

## Справочники
Обязательный столбец `code`. Должен быть первым.
Пример:
```
code,text
1234,Info about sth
```

## Run
Run this command in project root:
```commandLine
docker-compose up -d
```

## Testing

- После запуска инфраструктуры, можно зайти по ссылке:
curl -sSLk -XGET 'http://<Instanse_IP>:8585/documents?locale=RUS&type=patient.genders' | jq .
```
Иногда по непонятным причинам коннект к приложению зависает. При этом контейнеры запущены. Помогает ручной перезапуск виртуалки.
```

### By `go test`
To run all tests recursively, run this in project root:
```commandLine
$ docker-compose run --rm --name thesaurus_test thesaurus sh -c 'CGO_ENABLED=0 go test -cover -covermode atomic ./...'
```

## Разворачивание инфраструктуры.
- Для разворачивания инфраструктуры используется связка Packer+Terraform. В качестве платформы выбран Google Cloud Platform.
- Инфраструктурный репозиторий https://gitlab.com/project2019-02/infra 
- В первую очередь создаётся базовый образ для всех серверов. Операционная система CentOS 7. Файлы для Packer находятся в папке `packer`. Шаблон базового образа в файле `centos7-base.json`. Переменные объявлены в вайле `variables.json`. Команда для запуска сборки базового образа `packer build -var-file=variables.json centos7-base.json`.
- Далее посредством Terraform разворачиваются продуктовые сервера: сервер приложения с базой данных и сервер мониторинга. Сетевая инфраструктура также разворачивается посредством terraform. Файлы для Terraform находятся в папке `terraform`. Файлы с описанием сервера приложений и базы данных находятся в папке `modules/app`. Файлы с описанием сервера мониторинга находятся в папке `modules/mon`. Файлы с описанием сетевой инфраструктуры находятся в папке `modules/vpc`.
