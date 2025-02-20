run :
	@go run ./cmd/server/main.go

test:
	@go test -v ../...

migrate:
	@migrate create -ext sql -dir ./internal/database/postgres/migrations $(name)