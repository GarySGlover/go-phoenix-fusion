export GOROOT=

build:
	@tinygo build -target=pico -size short main.go

test:
	@go test -cover  ./pkg/...


flash : test build
	@tinygo flash -target=pico

release:
	@tinygo build -target=pico -no-debug

lint:
	@golangci-lint run pkg/... --fix
