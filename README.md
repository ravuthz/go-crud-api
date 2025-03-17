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

[![Open in IDX](https://cdn.idx.dev/btn/open_dark_32.svg)](https://idx.google.com/import?url=https://github.com/ravuthz/go-crud-api)

[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/github/ravuthz/go-crud-api)

[![Open in Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/ravuthz/go-crud-api)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/ravuthz/go-crud-api)

