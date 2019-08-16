# PG Gateway
<a href="https://goreportcard.com/report/github.com/just1689/pg-gateway"><img src="https://goreportcard.com/badge/github.com/just1689/pg-gateway" /></a><br />

<img align="right" height="200" src="docs/pg.png" />

This project aims to make it simple and fast to interact with a postgresql database over http.


Currently, the application supports inserts, get whole table, get row by id, get rows where field=value. 

## Setup

### Providing database connection details
| Env | Description |
|---|---|
| listenAddr | Address to listen on for http request (for ex: :8080) |
| pghost | Postgresql hostname |
| pguser | Postgresql username |
| pgpassword | Postgresql password |
| pgdb | Postgresql database |
| pgport | Postgresql port |
| poolSize | Number of postgresql connections to keep open |

- Build from source or
- Run the Docker container `docker pull just1689/scraping-in-go:svc-db-gateway`

You may need to set the search_path for the user.
`ALTER USER postgres SET search_path to myschema;`

## Usage
For the example below, we'll assume the application is hosted on localhost:8080

### Get all rows for that table
`curl http://localhost:8080/users`

### Get rows a table where x=y
`curl http://localhost:8080/users/x/y`

### Get row where id=z
`curl http://localhost:8080/users/z`

### Insert row into table
```
curl -X POST \
  http://localhost:8080/entities \
  -H 'Content-Type: application/json' \
  -d '{
	"entity": "user",
	"id": "12",
	"name": "Justin"
}'
```

### Memory Performance
Memory performance after 1M requests with a concurrency of 100. 
<img src="docs/memory3.png" />

### To do
- Fast http
- Websocket support (including client)
- Bulk insert
- Pagination
- Complex queries like postgrest
- Delete
- Update where

