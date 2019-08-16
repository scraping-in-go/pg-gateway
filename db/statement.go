package db

import (
	"fmt"
	"github.com/just1689/pg-gateway/model"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Insert(entity string, insertable model.Insertable) (err error) {
	conn := NextPoolCon()
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
