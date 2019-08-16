package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
	"strconv"
)

var DatabaseHost = os.Getenv("pghost")
var DatabaseUser = os.Getenv("pguser")
var DatabasePassword = os.Getenv("pgpassword")
var DatabaseDatabase = os.Getenv("pgdb")
var DatabasePort = os.Getenv("pgport")
var dp, _ = strconv.Atoi(DatabasePort)

func Connect() (conn *pgx.Conn, err error) {

	conn, err = pgx.Connect(pgx.ConnConfig{
		Host:     DatabaseHost,
		Port:     uint16(dp),
		User:     DatabaseUser,
		Password: DatabasePassword,
		Database: DatabaseDatabase,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	return

}
