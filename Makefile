.PHONY: build test

build:
	go build -i -o artifacts/compete github.com/LukeJelly/Cell-Compete-Go

test:
	go test -race $(go list ./... | grep -v vendor)
