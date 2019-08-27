# PostgreSQL Gateway
<a href="https://github.com/just1689/pg-gateway/releases"><img src="https://img.shields.io/badge/version-2.0-blue" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/just1689/pg-gateway"><img src="https://goreportcard.com/badge/github.com/just1689/pg-gateway"></a>&nbsp;<a href="https://codebeat.co/projects/github-com-just1689-pg-gateway-master"><img alt="codebeat badge" src="https://codebeat.co/badges/41278d9d-5877-4f6b-8638-9eec74b9aeba" /></a>&nbsp;<a href="https://codeclimate.com/github/just1689/pg-gateway/maintainability"><img src="https://api.codeclimate.com/v1/badges/28264ac4d5c087d43928/maintainability" /></a>&nbsp;[![CircleCI](https://circleci.com/gh/just1689/pg-gateway.svg?style=svg)](https://circleci.com/gh/just1689/pg-gateway)
<br />

<img align="right" height="240" src="docs/pg2.png" />

This project aims to make easy to interact with a PostgreSQL database over http.


## Features

Currently, the application supports the following database interactions 
- inserts, 
- update where field=value,
- delete where field=value,
- get whole table, 
- get row by id, 
- get with multiple field{comparison}value
- get but limit results
- get but select columns

Other features
- For multi-row returns, the application writes rows back to the client as each row is read.
- Low memory requirement (around 8 MB of RAM used for entire docker container under high load. Allocate 16 MB to be generous).
- Database connection cache. Each request doesn't have to wait for a new connection to be made.
- Transfer binary-to-binary. Reading from PostgreSQL is done in binary and written back to the client without conversion etc.
- Set the database connection details using environment variables. Great for Cloud Native environments. 
- Dockefile with light, low attack-surface final image.
- Listen address can be set by an environment variable.

## Setup

### Providing database connection details
| Env | Description |
|---|---|
| listenAddr | Address to listen on for http request (for ex: :8080) |
| pghost | PostgreSQL hostname |
| pguser | PostgreSQL username |
| pgpassword | PostgreSQL password |
| pgdb | PostgreSQL database |
| pgport | PostgreSQL port |
| poolSize | Number of PostgreSQL connections to keep open |

- Build from source or
- Run the Docker container `docker pull just1689/pg-gateway:latest`

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
is the equivalent of `SELECT * FROM users`

### Get n rows for that table
```shell script
curl http://localhost:8080/users?limit=10
```
is the equivalent of `SELECT * FROM users limit 10`


### Get particular columns for that table
```shell script
curl http://localhost:8080/users?select=id,name,email
```
is the equivalent of `SELECT id, name, email FROM users`


### Get rows a table where x=y
```shell script
curl http://localhost:8080/users?x=eq.y
```
is the equivalent of `SELECT * FROM users WHERE x=$1`



### Get rows where a=b and c>d
```shell script
curl http://localhost:8080/users?a=eq.b&c=gt.d
```
is the equivalent of `SELECT * FROM users WHERE a=$1 AND c>$2`

### Supported Comparisons
| Comparator | Explanation |
|---|---|
| eq | equals |
| lt | Less than |
| gt | Greater than |
| lte | Less than or equal to |
| gte | Greater than or equal to |
| neq | Not equal to |
| is | is for true, false |

### Get with all features!
```shell script
curl http://localhost:8080/users?a=eq.b&c=gt.d&select=email&limit=5
```
is the equivalent of `SELECT email FROM users WHERE a=$1 AND c>$2 limit 5`





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
is the equivalent of `INSERT INTO entities (entity, id, name) VALUES ("user", "12", "Justin")`



### Update rows where field=value
```shell script
curl -X PATCH http://localhost:8080/users/id/12
  http://localhost:8080/entities \
  -H 'Content-Type: application/json' \
  -d '{
	"entity": "user",
	"id": "12",
	"name": "Justin"
}'
```
is the equivalent of `update entities set entity=$1, id=$2, v=$3 where id=$4`



### Delete rows where field=value
```shell script
curl -X DELETE http://localhost:8080/users/id/12
```
is the equivalent of `DELETE FROM entities WHERE id=12`


### Understanding the response
An HTTP 200 OK is returned if everything went as expected.
 
Bad request is used if the request cannot be processed.

Internal Server Error is used for unexpected errors on the server-side.

Results for GET operations will always be contained in a JSON array: `[]` (even if expect to return one row). 


## Client

See the /client directory for the client. Examples of usage can be found in /examples.




### Memory Usage
Memory usage after 1 million requests (concurrency of 100 connections). 
<img src="docs/memory3.png" />


### Roadmap
| To do | Notes |
|---|---|
| Distinct | Allow distinct select for column. |
| Insert | Revisit implementation. |
| Update | Revisit implementation. |
| Delete | Revisit implementation. |
| Security | A strategy that makes sense for the context. |
| FastHTTP | Fewer allocs for each request. |
| Bulk inserts | Better insert performance. |
| Docker Service | Simple script to spin up db and service. |
| Benchmark | Read, write, allocs, memory performance. |
| Viper lib | 12 factor app style config |
| Websocket support | Great for less connection overhead |
| GRPC support | Investigation |

## Acknowledgments

This project was inspired by <a href="https://github.com/PostgREST/postgrest">PostgREST</a>. 
While PostgREST is specifically for rest over http, this project aims to enable less connection overhead by websocket and perhaps GRPC support. 