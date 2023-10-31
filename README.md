# Команды #
## Создать бд ## 
docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres
## Произвести миграцию up ##
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up
## Произвести миграцию down ##
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' down

# Примеры запросов #
## Запрос на регистрацию пользователя ##
### POST url - http://localhost:8000/auth/sign-up ###
{
    "name": "name",
    "username": "username",
    "password": "password"
}

## Запрос на авторизацию пользователя ##
### POST url - http://localhost:8000/auth/sign-in ###
{
    "username": "username",
    "password": "password"
}