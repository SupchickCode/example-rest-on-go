.SILENT:

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

postgres:
	docker run --name postgres -e POSTGRES_PASSWORD="qwerty" -p 5432:5432 -v pgdata:/var/lib/postgresql/data -d postgres