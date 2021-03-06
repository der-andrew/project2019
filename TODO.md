# TODO

## Требования
1. ~~Автоматизированные процессы создания и управления платформой~~
• ~~Ресурсы GCP~~
• ~~Инфраструктура для CI/CD~~
• ~~Инфраструктура для сбора обратной связи~~
2. ~~Использование практики IaC (Infrastructure as Code) для управления конфигурацией и инфраструктурой~~
3. ~~Настроен процесс CI/CD~~
4. ~~Все, что имеет отношение к проекту хранится в Git~~
5. ~~Настроен процесс сбора обратной связи~~
• ~~Мониторинг (сбор метрик, алертинг, визуализация)~~
• Логирование (опционально)
• Трейсинг (опционально)
• ChatOps (опционально)
6. Документация
• ~~README по работе с репозиторием~~
• ~~Описание приложения и его архитектуры~~
• ~~How to start?~~
• CHANGELOG с описанием выполненной работы
• Если работа в группе, то пункт включает автора изменений
7. Для сдачи проекта необходимы
• ссылка на Git-репозиторий проекта, с описанием по требованиям
• рекомендуем в рамках MVP или после сделать screenscast с рассказом о вашем проекте и его текущей реализации


## План работ.
- ~~Зарегиться на gitlab.com. Проверить возможности free версии. Если не устроит, развернуть gitlab в кубере (см. лекции и ДЗ), не забыв о DNS к нему для сборки приложения.~~
* Free версия вполне подходит.

- ~~Взять своё приложение, т.к. на ОТУС-ово много ругани в тредах(((.~~
1. https://otus-devops.slack.com/archives/CGJ7067K6/p1561802025411900
2. https://otus-devops.slack.com/archives/CGJ7067K6/p1561988909431600?thread_ts=1561966929.422100&cid=CGJ7067K6

- ~~Из приложения выпилить справочники.~~
- ~~Изучить приложение.~~
-- ~~Выпилить из приложения лишнее).~~
* ~~Swagger.~~
* ~~Pages.~~
* ~~test:postman~~
* ~~test:go_test~~
 
- GCP. ~~Работы ведём в проекте `docker-239319`~~.
- ~~CI/CD GitLab. Создать группу `Project2019-02` и проект в ней `thesaurus`.~~
- ~~CI/CD GitLab. В группе `Project2019-02` создать проект `infra`.~~
- ~~Прикрутить ssh ключи к проектам `thesaurus` `infra` с правами на запись.~~
```commandLine
ssh-keygen -t rsa -C 'for GitLab' -f ~/.ssh/GitLab
```
- ~~Подключить к проектам раннер(ы).~~ Не потребовалось.
- ~~IaC.~~
* ~~Базовый образ. Packer. Family: `common-base`.~~
* ~~Боевые машины. Terraform+Packer+Ansible.~~ ~~Molecule (may be for CI/CD testing) **NO!!! It stuff needs pip virtualenv!!! NO!!!**~~.
1. ~~app-host. Содержит приложение и базу данных.~~
2. ~~mon-host. Система мониторинга.~~

- ~~Monitoring.~~
* ~~app-host:~~
1. ~~cadvisor~~
2. ~~node-exporter~~
3. ~~mongodb_exporter~~

* ~~mon-host:~~
1. ~~prometheus~~
2. ~~grafana~~
~~dashboards:~~
~~cAdvisor+Node-Exporter~~ - https://grafana.com/dashboards/893, https://grafana.com/dashboards/395
~~MongoDB~~ - https://grafana.com/dashboards/5270
3. ~~alertmanager~~
4. ~~cadvisor~~
5. ~~node-exporter~~

- ~~Alerting.~~

## Замечания куратора
- ~~Управление инфраструктурой и конфигурацией. Преимущества~~
- ~~Terraform для разворота инфраструктуры. Рекомендации: Подробно описать как происходит разворачивание инфраструктуры.~~ Составить логическую схему взаимодействия компонентов.
- CI\CD
    Преимущества
    Рекомендации
Актуализировать пайплайн, убрать лишние закомментированные строки, составить логическую схему.
- Инструменты для работы с обратной связью
    Преимущества
- Мониторинг. Алертинг
    Рекомендации
Добавить инструкцию с описанием мониторинга. (Не увидел в общей репе)

Записать скринкаст, незнакомое приложение.
Краткий отзыв и мои хотелки
