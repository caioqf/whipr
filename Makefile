.PHONY: build

build:
	go build -o build/whipr

run: 
	make build && ./build/whipr