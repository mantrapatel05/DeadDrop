.PHONY: build run test clean tidy vet all

build:
	go build -o deaddrop.exe .

run:
	go run .

test:
	go test ./...

clean:
	rm -f deaddrop.exe
	rm -f server.exe

tidy:
	go mod tidy

vet:
	go vet ./...

all: vet tidy build