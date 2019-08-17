# PG Gateway
<a href="https://github.com/just1689/pg-gateway/releases"><img src="https://img.shields.io/badge/version-1.0-blue" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/just1689/pg-gateway"><img src="https://goreportcard.com/badge/github.com/just1689/pg-gateway" /></a><br />

<img align="right" height="200" src="docs/pg2.png" />

This project aims to make it simple and fast to interact with a postgresql database over http inspired by <a href="https://github.com/PostgREST/postgrest">Postgrest</a>.


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
```sql
ALTER USER postgres SET search_path to myschema
```

## Usage
For the examples below, we'll assume the application is hosted on localhost:8080

### Get all rows for that table
```shell script
curl http://localhost:8080/users
```
is the equivalent of  
```sql
SELECT * FROM users
```

### Get rows a table where x=y
```shell script
curl http://localhost:8080/users/x/y
```
is the equivalent of  
```sql
SELECT * FROM users WHERE x=y
```

### Get row where id=z
```shell script
curl http://localhost:8080/users/z
```
is the equivalent of  
```sql
SELECT * FROM users WHERE id=z
```

### Insert row into table
```shell script
curl -X POST \
  http://localhost:8080/entities \
  -H 'Content-Type: application/json' \
  -d '{
	"entity": "user",
	"id": "12",
	"name": "Justin"
}'
```
is the equivalent of  
```sql
INSERT INTO entities (entity, id, name) VALUES ("user", "12", "Justin")
```


### Memory Usage
Memory performance after 1M requests with a concurrency of 100. 
<img src="docs/memory3.png" />

### Roadmap
| To do | Notes |
|---|---|
| Deleting by id | Important feature. |
| Better standard for returning errors | Important feature. |
| FastHTTP | Fewer allocs for each request. |
| Bulk insert | Better insert performance. |
| DockerHub | Project, repository and notes. |
| Docker Service | Simple script to spin up db and service. |
| Pagination | Restricting result sets. |
| Complex queries | HUGE amount of work. |
| Deleting (complex) | Requires that complex queries are implemented. |
| Updating rows | Requires that complex queries are implemented. |

