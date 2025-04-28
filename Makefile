run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go
test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
lint:
	golangci-lint run ./...

db:
	docker-compose up --build

down:
	docker-compose down