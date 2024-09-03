## Запуск локально

### Создать `.env.dev`

```bash
touch task_tracker_app/.env.dev
```

#### Пример переменных окружения
```
DB_NAME=golang_to_do_list
DB_HOST=db
DB_PORT=5432
DB_USER=example_user
DB_PASS=example_pass
DB_SSL_MODE=disable

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
```
### Установливаем библиотеку для миграций

```bash
go get -u github.com/golang-migrate/migrate/v4
```
### Запускаем сервер
```bash
go run main.go
```

### Прогоняем миграции
```bash
make migrate_up
```



P.S. Данный проект основан на моей предыдущем проекте: https://github.com/JustArtur/golang_task_tracker.
Убрано лишнее и сделано все под Ваше ТЗ.
