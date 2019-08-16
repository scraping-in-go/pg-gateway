# PG Gateway
<img align="right" width="290" height="290" src="docs/pg.png" />



This project aims to make it simple and fast to interact with a postgresql database over http.

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
| poolSize | Number of postgresql connection to keep open |

- Build from source or
- Run the Docker container `docker pull just1689/scraping-in-go:svc-db-gateway`

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
  -H 'Postman-Token: da996145-35d6-4e19-88b6-f7395cb229ce' \
  -H 'cache-control: no-cache' \
  -d '{
	"entity": "user",
	"id": "12",
	"name": "Justin"
}'
```


### To do
- Fast http
- Complex queries like postgrest
- Delete
- Update where
- Bulk insert
- Pagination
- Websocket support (including client)

