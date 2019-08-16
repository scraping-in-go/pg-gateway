package db

import (
	"fmt"
	"github.com/scraping-in-go/svc-db-gateway/model"
	"github.com/sirupsen/logrus"
	"strconv"
)

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
