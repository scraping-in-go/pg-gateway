package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

func GetEntityAll(entity string) (result chan []byte, err error) {
	result = make(chan []byte)
	conn := NextPoolCon()
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s", entity, entity)
	rows, err := conn.Query(sql)
	if err != nil {
		logrus.Errorln(err)
		conn.Close()
		return
	}
	go rowsToChan(rows, result, func() { conn.Close() })
	return
}

func GetEntityMany(entity, field, id string) (result chan []byte, err error) {
	result = make(chan []byte)
	conn := NextPoolCon()
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s where %s=$1", entity, entity, field)
	rows, err := conn.Query(sql, id)
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

func GetEntityByID(entity, id string) (row []byte, err error) {
	conn := NextPoolCon()
	defer conn.Close()
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s where id=$1", entity, entity)
	if err = conn.QueryRow(sql, id).Scan(&row); err != nil {
		logrus.Errorln(err)
		return
	}
	return

}
