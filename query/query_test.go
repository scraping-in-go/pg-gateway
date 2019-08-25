package query

import (
	"fmt"
	"testing"
)

func TestQuery_ToQuerySimple(t *testing.T) {
	query := Query{
		Entity: "users",
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "100",
			},
		},
		Limit: 1000,
	}
	expected := "SELECT * FROM users WHERE id=$1 LIMIT 1000"
	sql, binds := query.ToQuery()
	if sql != expected {
		t.Error("bad sql generation, found:")
		t.Error(sql)
		t.Error(expected)
		for _, row := range binds {
			fmt.Println(row)
		}
		return
	}

}

func TestQuery_ToQuerySimple2(t *testing.T) {
	query := Query{
		Entity: "users",
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "100",
			},
			{
				Field:      "age",
				Comparator: "gte",
				Value:      "50",
			},
		},
		Limit: 1000,
	}
	expected := "SELECT * FROM users WHERE id=$1 AND age>$2 LIMIT 1000"
	sql, binds := query.ToQuery()
	if sql != expected {
		t.Error("bad sql generation, found:")
		t.Error(sql)
		t.Error(expected)
		for _, row := range binds {
			fmt.Println(row)
		}
		return
	}

}

func TestQuery_ToQuerySimpler(t *testing.T) {
	query := Query{
		Entity: "users",
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "100",
			},
		},
	}
	expected := "SELECT * FROM users WHERE id=$1"
	sql, binds := query.ToQuery()
	if sql != expected {
		t.Error("bad sql generation, found:")
		t.Error(sql)
		t.Error(expected)
		for _, row := range binds {
			fmt.Println(row)
		}
		return
	}

}
