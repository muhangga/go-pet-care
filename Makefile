run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go

test:
	go test -v ./...

migrate:
	migrate -path config/db/migrations -database "postgres://postgres:root@localhost:5432/go-pet?sslmode=disable" -verbose up
	
migrate-down:
	migrate -path config/db/migrations -database "postgres://postgres:root@localhost:5432/go-pet?sslmode=disable" -verbose down