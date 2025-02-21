run :
	@docker compose build && docker compose up

test:
	@go test -v ../...

migrate:
	@migrate create -ext sql -dir ./internal/database/postgres/migrations $(name)