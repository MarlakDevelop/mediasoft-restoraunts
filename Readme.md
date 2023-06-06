# Краткое описание проекта и навигация по нему
## Проект запускается через docker-compose
```shell
docker-compose up --build
```
## Переменные проекта(.env)
- __SERVICE_MONITORING_PORT__ - Порт панели мониторинга сервисов(пока бесполезна, так как является заделом на будущее расширение)
- __CUSTOMER_OFFICE_SERVICE_HOST__ - Хост "Пользовательского" сервиса
- __CUSTOMER_OFFICE_SERVICE_PORT__ - Порт "Пользовательского" сервиса(GRPC)
- __RESTAURANT_MENU_SERVICE_HOST__ - Хост "Ресторанного" сервиса
- __RESTAURANT_MENU_SERVICE_PORT__ - Порт "Ресторанного" сервиса(GRPC)
- __STATISTICS_STATISTICS_SERVICE_HOST__ - Хост "Статистического" сервиса
- __STATISTICS_STATISTICS_SERVICE_PORT__ - Порт "Статистического" сервиса(GRPC)
- __CUSTOMER_OFFICE_SERVICE_HTTP_GATEWAY_PORT__ - Порт "Пользовательского" сервиса(HTTP/JSON)
- __RESTAURANT_MENU_SERVICE_HTTP_GATEWAY_PORT__ - Порт "Ресторанного" сервиса(HTTP/JSON)
- __STATISTICS_STATISTICS_SERVICE_HTTP_GATEWAY_PORT__ - Порт "Статистического" сервиса(HTTP/JSON)
- __DB_MONITORING_PORT__ - Порт панели мониторинга баз данных(pgadmin)
- __DB_MONITORING_EMAIL__ - Почта от панели мониторинга баз данных(pgadmin)
- __DB_MONITORING_PASSWORD__ - Пароль от панели мониторинга баз данных(pgadmin)
- __CUSTOMER_OFFICE_DB_PORT__ - Порт сервера базы данных "Пользовательского" сервиса
- __CUSTOMER_OFFICE_DB_HOST__ - Хост сервера базы данных "Пользовательского" сервиса
- __CUSTOMER_OFFICE_DB_NAME__ - Название базы данных "Пользовательского" сервиса
- __CUSTOMER_OFFICE_DB_USERNAME__ - Имя пользователя от сервера базы данных "Пользовательского" сервиса
- __CUSTOMER_OFFICE_DB_PASSWORD__ - Пароль от сервера базы данных "Пользовательского" сервиса
- __RESTAURANT_MENU_DB_PORT__ - Порт сервера базы данных "Ресторанного" сервиса
- __RESTAURANT_MENU_DB_HOST__ - Хост сервера базы данных "Ресторанного" сервиса
- __RESTAURANT_MENU_DB_NAME__ - Название базы данных "Ресторанного" сервиса
- __RESTAURANT_MENU_DB_USERNAME__ - Имя пользователя от сервера базы данных "Ресторанного" сервиса
- __RESTAURANT_MENU_DB_PASSWORD__ - Пароль от сервера базы данных "Ресторанного" сервиса
- __STATISTICS_STATISTICS_DB_PORT__ - Порт сервера базы данных "Статистического" сервиса
- __STATISTICS_STATISTICS_DB_HOST__ - Хост сервера базы данных "Статистического" сервиса
- __STATISTICS_STATISTICS_DB_NAME__ - Название базы данных "Статистического" сервиса
- __STATISTICS_STATISTICS_DB_USERNAME__ - Имя пользователя от сервера базы данных "Статистического" сервиса
- __STATISTICS_STATISTICS_DB_PASSWORD__ - Пароль от сервера базы данных "Статистического" сервиса
- __AMQP_MONITORING_PORT__ - Порт панели мониторинга брокера сообщений(RabbitMq)
- __AMQP_QUEUING_HOST__ - Хост брокера сообщений(RabbitMq)
- __AMQP_QUEUING_PORT__ - Порт брокера сообщений(RabbitMq)
- __AMQP_USERNAME__ - Имя пользователя брокера сообщений(RabbitMq)
- __AMQP_PASSWORD__ - Пароль брокера сообщений(RabbitMq)
- __AMQP_ORDER_RESTAURANT_QUEUE_NAME__ - Название очереди сообщений между "Пользовательским" и "Ресторанным" сервисами для получения заказов
- __AMQP_ORDER_STATISTICS_QUEUE_NAME__ - Название очереди сообщений между "Пользовательским" и "Статистическим" сервисами для получения заказов

### Для проверки GRPC рекомендую Postman