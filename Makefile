# ==== sqlboiler generate ====
.PHONY: models
models:
	sqlboiler mysql

# ==== database migration (golang-migrate) ====
.PHONY: migrate
migrate:
	migrate -path schema -database "mysql://root:rootpass@tcp(localhost:3306)/db" up

.PHONY: app
app:
	go run main.go