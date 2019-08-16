# PG Gateway

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
### Get all rows for that table

- Get rows for key=value for table
- Get row where id=X
- Insert row into table

### To do
- Complex queries like postgrest
- Delete
- Update where
- Bulk insert
- Pagination
- Websocket support (including client)

