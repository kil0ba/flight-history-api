default: start

# Migrate - либа для миграций бд
# https://github.com/golang-migrate/migrate
.PHONY: install-migrate
install-migrate-windows:
	scoop install migrate

.PHONY: install-migrate-mac
install-migrate-mac:
	brew install golang-migrate

.PHONY: compose-up
compose-up:
	podman compose --compatibility up --build

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate-db
migrate-db:
	migrate -path migrations -database "postgres://localhost/flighthistory?sslmode=disable&user=root&password=example" up

.PHONY: migrate-db
migrate-db-down:
	migrate -path migrations -database "postgres://localhost/flighthistory?sslmode=disable&user=root&password=example" down

.PHONY: start
start:
	go run ./cmd/flight-history-api/main.go