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
			co.Comparator = riArr[0]
			co.Value = riArr[1]
			result.Comparisons = append(result.Comparisons, co)
		}
	}
	return
}
