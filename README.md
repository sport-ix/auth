docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up