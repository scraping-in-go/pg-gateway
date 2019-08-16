package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
	"strconv"
)

var DatabaseHost = os.Getenv("pghost")
var DatabaseUser = "postgres"
var DatabasePassword = os.Getenv("pgpassword")
var DatabaseDatabase = "scrapedb"
var DatabasePort = os.Getenv("pgport")

func Connect() (conn *pgx.Conn, err error) {
	dp, _ := strconv.Atoi(DatabasePort)

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
