package query

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Query struct {
	Entity      string
	Select      []string
	Comparisons []Comparison
	Limit       int
}

/*
select
   (select row_to_json(_) as row from (select tbl.schemaname, tbl.tablename) as _) as schemaname
from
   pg_tables as tbl
*/

func (q *Query) ToQuery() (queryString string, bindArr []string) {
	queryString = "SELECT "
	if len(q.Select) == 0 {
		queryString += "row_to_json($0) as row"
		bindArr = append(bindArr, q.Entity)
	} else {
		queryString += "(select row_to_json(_) as row from (select tbl."
		for id, row := range q.Select {
			if id != 0 {
				queryString += ", tbl."
			}
			queryString += "$" + strconv.Itoa(len(bindArr))
			bindArr = append(bindArr, row)
		}
		queryString += ") as _) as schemaname"

	}

	queryString += " FROM $1 as tbl"
	if len(q.Comparisons) != 0 {
		queryString += " WHERE "
	}
	bindArr = append(bindArr, q.Entity)
	for id, row := range q.Comparisons {
		if id > 0 {
			queryString += " AND "
		}
		queryString += row.Field + row.ComparatorToSQL() + "$" + strconv.Itoa(len(bindArr))
		bindArr = append(bindArr, row.Value)
	}
	if q.Limit != 0 {
		queryString += " LIMIT $" + strconv.Itoa(len(bindArr))
		bindArr = append(bindArr, strconv.Itoa(q.Limit))
	}
	fmt.Println(queryString)
	for _, row := range bindArr {
		fmt.Println(row)
	}
	return
}

func (q *Query) processOtherQuery(field, value string) {
	if field == "limit" {
		i, _ := strconv.Atoi(value)
		q.Limit = i
		return
	} else if field == "select" {
		q.Select = strings.Split(value, ",")
	}
	logrus.Errorln("not sure how to handle", field, value, "for processOtherQuery")

}

type Comparison struct {
	Field      string
	Comparator string
	Value      string
}

func (c *Comparison) ComparatorToSQL() string {
	if c.Comparator == "eq" {
		return "="
	}
	if c.Comparator == "gte" {
		return ">"
	}
	if c.Comparator == "lte" {
		return "<"
	}

	return c.Comparator
}
