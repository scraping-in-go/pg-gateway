package query

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var comparisonMap = map[string]string{
	"eq":  "=",
	"gt":  ">",
	"lt":  "<",
	"gte": ">=",
	"lte": "<=",
	"neg": "!=",
	"is":  "is",
}

type Query struct {
	Entity      string
	Select      []string
	Comparisons []Comparison
	Limit       int
}

func (q *Query) ToURL(baseURL string) string {
	url := baseURL + "/" + q.Entity
	changed := 0
	for _, sel := range q.Select {
		changed++
		if changed == 1 {
			url += "?select="
		} else {
			url += ","
		}
		url += sel
	}

	for _, c := range q.Comparisons {
		changed++
		if changed == 1 {
			url += "?"
		} else {
			url += "&"
		}
		url += c.ComparatorToURL()
	}

	if q.Limit != 0 {
		changed++
		if changed == 1 {
			url += "?"
		} else {
			url += "&"
		}
		url += "limit=" + strconv.Itoa(q.Limit)
	}
	return url

}

func (q *Query) ToQuery() (queryString string, bindArr []interface{}) {
	queryString = "SELECT "
	if len(q.Select) == 0 {
		queryString += "row_to_json(" + "tbl" + ") as row"
	} else {
		queryString += "(select row_to_json(_) as row from (select tbl."
		for id, row := range q.Select {
			if id != 0 {
				queryString += ", tbl."
			}
			queryString += row
		}
		queryString += ") as _) as schemaname"

	}

	queryString += " FROM " + q.Entity + " as tbl"
	if len(q.Comparisons) != 0 {
		queryString += " WHERE "
	}
	for id, row := range q.Comparisons {
		if id > 0 {
			queryString += " AND "
		}
		queryString += row.Field + row.ComparatorToSQL() + "$" + strconv.Itoa(len(bindArr)+1)
		bindArr = append(bindArr, row.Value)
	}
	if q.Limit != 0 {
		queryString += " LIMIT $" + strconv.Itoa(len(bindArr)+1)
		bindArr = append(bindArr, strconv.Itoa(q.Limit))
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
		return
	}
	//TODO: determine whether or not this error should be fatal. If you subscribe to the idea
	// that tests should cover possible use cases, it may be useful to panic here.
	logrus.Errorln("not sure how to handle", field, value, "for processOtherQuery")

}

type Comparison struct {
	Field      string
	Comparator string
	Value      string
}

func (c *Comparison) ComparatorToURL() string {
	return c.Field + "=" + c.Comparator + "." + c.Value
}
func (c *Comparison) ComparatorToSQL() string {
	v, found := comparisonMap[c.Comparator]
	if !found {
		return c.Comparator
	}
	return v

}
