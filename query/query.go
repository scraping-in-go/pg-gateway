package query

import (
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

func (q *Query) ToQuery() (queryString string, bindArr []string) {
	queryString = "SELECT "
	if len(q.Select) == 0 {
		queryString += "*"
	} else {
		for id, s := range q.Select {
			if id != 0 {
				queryString += ", "
			}
			queryString += s
		}
	}

	queryString += " FROM " + q.Entity + " WHERE "
	for id, row := range q.Comparisons {
		if id > 0 {
			queryString += " AND "
		}
		queryString += row.Field + row.ComparatorToSQL() + "$" + strconv.Itoa(id+1)
		bindArr = append(bindArr, row.Value)
	}
	if q.Limit != 0 {
		queryString += " LIMIT " + strconv.Itoa(q.Limit)
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
