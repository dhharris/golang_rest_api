# REST API exercise in Go

## Dependencies
* [MySQL Community Server](https://dev.mysql.com/downloads/mysql/)
  * MySQL credentials are hard-coded. Change them as-needed in the code
* [Go](https://go.dev/doc/install) Programming language

## Before starting the server
* Create the database in MySQL using the source file in `sql/schema.sql`
* Install Go modules using `go get ./web-server`


## Starting the server
```
go run .
```

## Example queries
Creating a new user
```
curl http://localhost:8080/user --include --header "Content-Type: application/json" --request "POST" --data '{"name": "test"}'
```
This will return a `uuid` that may be used in the following queries

Get game state information
```
curl http://localhost:8080/user/<uuid>/state --include --header "Content-Type: application/json" --request "GET"
```

