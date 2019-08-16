package db

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

func StartPool(size int) func() *pgx.Conn {
	pool := make(chan *pgx.Conn, size)
	go func() {
		for {
			c, err := Connect()
			if err != nil {
				logrus.Panic(err)
			}
			pool <- c
		}
	}()
	return func() *pgx.Conn {
		c := <-pool
		return c
	}
}
