build:
	@go build -o bin/blog ./cmd/blog/main.go

run: build
	@go run /bin/blog