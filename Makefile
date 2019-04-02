build:
	go build -o meme cmd/meme/main.go cmd/meme/root.go cmd/meme/build.go

build-repl:
	go build -o meme-repl cmd/repl/main.go
	./meme-repl

test:
	ginkgo ./...
