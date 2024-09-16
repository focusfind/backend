build:
	go build -o bin/focusfind_server ./cmd

run: build
	./bin/focusfind_server

test:
	go test -v ./...
