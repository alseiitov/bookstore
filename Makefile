.PHONY: build migrate

build:
	CGO_ENABLED=0 go build -o ./build/api ./cmd/main.go

migrate:
	migrate -path ./migrations -database 'postgres://postgres:secret@0.0.0.0:5432/postgres?sslmode=disable' up