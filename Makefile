check:
	go vet -v ./...; \
	go fmt ./...

build:
	go build -v ./...

test:
	CGO_ENABLED=0 go test -v ./...

server:
	go run server.go

client:
	go run client.go

