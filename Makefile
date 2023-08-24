build:
	@go build -o bin/blog ./cmd/blog/main.go

run: build
	@./bin/blog
