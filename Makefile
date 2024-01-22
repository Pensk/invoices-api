.PHONY: shell

build:
	docker compose build

up:
	docker compose up

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
