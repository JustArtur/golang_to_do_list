include to_do_list_app/.env.dev

DB_URL ?= postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

MIGRATIONS_PATH = to_do_list_app/db/migrate

g_migration:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)

migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate_version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

lint:
	gofmt -l . && golint ./...