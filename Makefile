test:
	@go test -v ./...

templ-generate:
	@templ generate

docker-pg-run:
	@docker-compose up pg-local

docker-stop:
	@docker-compose down

run: templ-generate
	@go run cmd/main.go