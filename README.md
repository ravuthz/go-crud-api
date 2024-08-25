```bash

cd ~/Projects/golang

mkdir go-crud-api && cd go-crud-api

go mod init go-crud-api

mkdir db router
touch main.go .env

go install github.com/gin-gonic/gin@latest
go install github.com/jinzhu/gorm@latest
go install github.com/google/uuid@latest
go install github.com/joho/godotenv@latest

go mod tidy

go run main.go

```

```bash

curl -L 'localhost:9000/movies' -H 'Content-Type: application/json' -d '{
    "name": "movie-1",
    "description": "movie-1"
}'

curl -L 'localhost:9000/movies' -H 'Content-Type: application/json' -d '{
    "name": "movie-2",
    "description": "movie-2"
}'

curl -L 'localhost:9000/movies' -H 'Content-Type: application/json' -d '{
    "name": "movie-3",
    "description": "movie-3"
}'

curl -L -X GET 'http://localhost:9000/movies' | jq

```
