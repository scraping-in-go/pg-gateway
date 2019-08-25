package query

import (
	"strings"
)

type Query struct {
	Entity      string
	Comparisons []Comparison
}

type Comparison struct {
	Field      string
	Comparator string
	Value      string
}

func BuildQueryFromURL(u string) (result Query) {
	result = Query{
		Comparisons: []Comparison{},
	}
	u = u[1:]
	idx := strings.Index(u, "?")
	if idx == -1 {
		return
	}

	result.setEntityFromUrl(u, idx)

	//Remove the query from the url
	u = u[idx+1:]

	//Build a comparison
	co := Comparison{}

	idx = strings.Index(u, "=")
	co.Field = u[0:idx]
	u = u[idx+1:]

	idx = strings.Index(u, ".")
	co.Comparator = u[0:idx]
	u = u[idx+1:]

	if strings.Index(u, "&") == -1 {
		co.Value = u[0:idx]
	} else {
		co.Value = u[0 : idx-1]
		u = u[idx-1:]
	}

	result.Comparisons = append(result.Comparisons, co)

	another := strings.Index(u, "&") != -1
	u = u[1:]
	for another {

		//Build a comparison
		co := Comparison{}

		idx = strings.Index(u, "=")
		co.Field = u[0:idx]
		u = u[idx+1:]

		idx = strings.Index(u, ".")
		co.Comparator = u[0:idx]
		u = u[idx+1:]

		if strings.Index(u, "&") == -1 {
			co.Value = u
		} else {
			co.Value = u[0 : idx-1]
			u = u[idx-1:]
		}

		result.Comparisons = append(result.Comparisons, co)

		another = strings.Index(u, "&") != -1
	}

	return

}

func (q *Query) setEntityFromUrl(url string, idOfQ int) {
	q.Entity = url[0:idOfQ]
}
