.PHONY: build

build:
	go build -o build/whipr

build-bin:
	go build -o ~/.local/bin/whipr

run: 
	make build && ./build/whipr