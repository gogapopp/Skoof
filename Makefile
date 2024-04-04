test:
	@go test -v ./...

templ-generate:
	@templ generate

run: templ-generate
	@go run cmd/main.go