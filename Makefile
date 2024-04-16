test:
	@go test -v ./...

templ-generate:
	@templ generate

docker-run:
	@docker-compose up

docker-stop:
	@docker-compose down

run: templ-generate
	@go run cmd/main.go