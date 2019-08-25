package db

import (
	"github.com/jackc/pgx"
	"github.com/just1689/pg-gateway/query"
	"github.com/sirupsen/logrus"
)

func GetByQuery(query query.Query) (result chan []byte, err error) {
	result = make(chan []byte)
	conn := NextPoolCon()
	sql, bind := query.ToQuery()
	rows, err := conn.Query(sql, bind)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln(sql)
		conn.Close()
		return
	}
	go rowsToChan(rows, result, func() { conn.Close() })
	return
}

func rowsToChan(rows *pgx.Rows, result chan []byte, closer func()) {
	for rows.Next() {
		s := []byte{}
		if err := rows.Scan(&s); err != nil {
			logrus.Errorln(err)
			continue
		}
		result <- s
	}
	close(result)
	rows.Close()
	closer()
}
