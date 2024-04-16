test:
	@go test -v ./...

templ-generate:
	@templ generate

stop:
	@docker-compose down

run: templ-generate
	@go run cmd/main.go