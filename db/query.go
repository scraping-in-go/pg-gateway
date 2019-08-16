package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var NextPoolCon = func() *pgx.Conn {
	return nil
}

func GetEntityAll(entity string) (result chan string, err error) {
	result = make(chan string)
	conn := NextPoolCon()
	if err != nil {
		logrus.Errorln(err)
		conn.Close()
		return
	}
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s", entity, entity)
	rows, err := conn.Query(sql)
	if err != nil {
		logrus.Errorln(err)
		conn.Close()
		return
	}
	go func() {
		defer conn.Close()
		defer close(result)
		for rows.Next() {
			s := ""
			err := rows.Scan(&s)
			if err != nil {
				logrus.Errorln(err)
			}
			result <- s
		}
		if err != nil {
			logrus.Errorln(err)
			return
		}
	}()
	return
}

func GetEntityMany(entity, field, id string) (result chan string, err error) {
	result = make(chan string)
	conn := NextPoolCon()
	if err != nil {
		logrus.Errorln(err)
		conn.Close()
		return
	}
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s where %s=$1", entity, entity, field)
	rows, err := conn.Query(sql, id)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln(sql)
		conn.Close()
		return
	}
	go func() {
		defer conn.Close()
		defer close(result)
		for rows.Next() {
			s := ""
			err := rows.Scan(&s)
			if err != nil {
				logrus.Errorln(err)
			}
			result <- s
		}
		if err != nil {
			logrus.Errorln(err)
			return
		}
	}()
	return
}

func GetEntityByID(entity, id string) (row string, err error) {
	conn := NextPoolCon()
	defer conn.Close()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	sql := fmt.Sprintf("select row_to_json(%s)as row from %s where id=$1", entity, entity)
	err = conn.QueryRow(sql, id).Scan(&row)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}
