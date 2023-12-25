build:
	docker-compose build crud-go

run:
	docker-compose up crud-go

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@172.18.0.1:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go