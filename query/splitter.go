package query

import "strings"

type Query struct {
	Entity string

	Field      string
	Comparator string
	Value      string
}

func BuildQueryFromURL(u string) (result Query) {
	result = Query{}
	u = u[1:]
	idx := strings.Index(u, "?")
	if idx == -1 {
		return
	}
	result.Entity = u[0:idx]
	u = u[idx+1:]

	idx = strings.Index(u, "=")
	result.Field = u[0:idx]
	u = u[idx+1:]

	idx = strings.Index(u, ".")
	result.Comparator = u[0:idx]
	u = u[idx+1:]

	result.Value = u[0:idx]

	return

}
