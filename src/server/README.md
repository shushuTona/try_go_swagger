# server

## install sqlc

> go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## generate code by swagger

> cd /go/src/server && swagger generate server -t ./gen -f ../swagger/swagger.yaml --exclude-main

## generate code by sqlc

> cd /go/src/server/sqlc && sqlc generate

## test post request

> curl -X POST -H "Content-Type: application/json" -d '{"Title":"create task1", "Desc":"task1 desc"}' localhost:8000/api/task
