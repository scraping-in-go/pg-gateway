package query

import (
	"strings"
)

func SplitURL(url string) []string {
	result := strings.Split(url, "?")
	if len(result) == 1 {
		return result
	}
	results := make([]string, 1)
	arr := strings.Split(result[1], "&")
	results[0] = result[0]
	results = append(results, arr...)
	return results

}
