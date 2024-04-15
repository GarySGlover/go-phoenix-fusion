export GOROOT=

build:
	@tinygo build -target=pico -size short main.go

test: build
	@tinygo test -target=pico

flash : test
	@tinygo flash -target=pico

release:
	@tinygo build -target=pico -no-debug
