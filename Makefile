build:
	go build cmd/meme/meme.go

build-repl:
	go build cmd/repl/main.go

test:
	ginkgo ./...
