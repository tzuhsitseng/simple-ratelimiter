install:
	@go mod download
	@go mod tidy

build: install
	@go build -o app github.com/gogotsenghsien/simple-rate-limit/src

run: build
	@./app