build:
	go build cmd/meme/meme.go

build-repl:
	go build cmd/repl/repl.go
	./repl

test:
	ginkgo ./...
