test:
	@go test -v ./...

templ:
	@templ generate

run: templ
	@go run cmd/main.go