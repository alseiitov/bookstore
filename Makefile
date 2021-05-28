.PHONY: build migrate

build:
	CGO_ENABLED=0 go build -o ./build/api ./cmd/main.go

migrate:
	migrate -path ./migrations -database 'postgres://postgres:secret@0.0.0.0:5432/postgres?sslmode=disable' up

run: build
	docker-compose up

clean:
	sudo rm -rf .database  
	docker image rm -f bookstore_bookstore-api
	docker container rm -f bookstore_bookstore-api_1 bookstore-db
	
swagger:
	swag init --parseDependency -g internal/app/app.go 