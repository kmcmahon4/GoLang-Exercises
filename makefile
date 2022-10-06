run:
	go run ./cmd/main.go

build:
	go build -o bin/main cmd/main.go
	
deps:
	go mod download

fmt: 
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet
	golangci-lint run

test:
	go test ./...