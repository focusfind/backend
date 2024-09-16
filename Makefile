build:
	go build -o bin/spotmeserver

run: build
	./bin/spotmeserver

test:
	go test -v ./...
