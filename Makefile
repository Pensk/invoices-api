.PHONY: shell setup

setup: build up-daemon mod-tidy migrate seed-db

build:
	docker compose build

up:
	docker compose up

up-daemon:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

shell:
	docker compose run --rm invoices bash

mysql:
	docker compose exec mysql mysql -u invoice -d invoices -p invoice

mod-tidy:
	docker compose run --rm invoices go mod tidy

mysql:
	docker compose exec mysql mysql -uinvoice -D invoices -pinvoice

migrate:
	docker compose run --rm migrate -path=/migrations/ -database="mysql://invoice:invoice@tcp(mysql:3306)/invoices" up

migrate-down:
	docker compose run --rm migrate -path=/migrations/ -database="mysql://invoice:invoice@tcp(mysql:3306)/invoices" down -all

seed-db:
	docker compose exec -T mysql mysql -uinvoice -D invoices -pinvoice < go/db/fixtures/company_data.sql

test:
	docker compose exec invoices go test -v ./...
