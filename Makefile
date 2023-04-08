.SILENT:

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

postgres:
	docker run -e POSTGRES_PASSWORD="qwerty" -p 5432:5432 -d postgres