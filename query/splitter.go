package query

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Query struct {
	Entity      string
	Comparisons []Comparison
	Limit       int
}

func (q *Query) processOtherQuery(field, value string) {
	if field == "limit" {
		i, _ := strconv.Atoi(value)
		q.Limit = i
		return
	}
	logrus.Fatal("not sure how to handle", field, value, "for processOtherQuery")

}

type Comparison struct {
	Field      string
	Comparator string
	Value      string
}

func BuildQueryFromURL(u string) (result Query) {
	split := SplitURL(u)
	result = Query{
		Entity:      split[0],
		Comparisons: []Comparison{},
	}
	if len(split) > 1 {
		for id, row := range split {
			if id == 0 {
				continue
			}
			coArr := strings.Split(row, "=")
			co := Comparison{
				Field: coArr[0],
			}
			riArr := strings.Split(coArr[1], ".")
			if len(riArr) == 1 {
				result.processOtherQuery(coArr[0], riArr[0])
				continue
			}
			co.Comparator = riArr[0]
			co.Value = riArr[1]
			result.Comparisons = append(result.Comparisons, co)
		}
	}
	return
}
