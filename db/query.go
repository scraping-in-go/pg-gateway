package db

import (
	"fmt"
	"github.com/scraping-in-go/svc-db-gateway/model"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetEntityAll(entity string) (result chan string, err error) {
	result = make(chan string)
	conn, err := Connect()
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
	conn, err := Connect()
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
	conn, err := Connect()
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

func Insert(entity string, insertable model.Insertable) (err error) {
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	cols := ""
	binds := ""
	colID := 0
	vals := make([]interface{}, 0)
	for colName, val := range insertable {
		va := string(val)
		if va[:1] == "\"" {
			va = va[1 : len(va)-1]
			vals = append(vals, va)
		} else {
			vals = append(vals, val)
		}
		colID++
		if cols == "" {
			cols = colName
		} else {
			cols += ", " + colName
		}
		if binds == "" {
			binds = "$" + strconv.Itoa(colID)
		} else {
			binds += ", $" + strconv.Itoa(colID)
		}
	}

	sql := fmt.Sprintf("insert into %s (%s) values(%s)", entity, cols, binds)
	_, err = conn.Exec(sql, vals...)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}
