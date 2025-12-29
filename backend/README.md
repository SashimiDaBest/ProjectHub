## APPLY UP MIGRATIONS
migrate -path db/migrations -database "postgres://myuser:mypassword@localhost:5432/myprojectdb?sslmode=disable" up

## APPLY DOWN MIGRATIONS 
migrate -path db/migrations -database "postgres://myuser:mypassword@localhost:5432/myprojectdb?sslmode=disable" down

[localhost:8000](https://localhost:8000)

## BASIC APP COMMANDS
go run main.go