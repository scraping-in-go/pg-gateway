package query

import (
	"fmt"
	"testing"
)

func TestQuery_ToQuery(t *testing.T) {
	query := Query{
		Entity:      "users",
		Comparisons: []Comparison{},
	}
	expected := "SELECT row_to_json($0) as row FROM $1 as tbl"
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
	expected := "SELECT row_to_json($0) as row FROM $1 as tbl WHERE id=$2 LIMIT $3"
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
	expected := "SELECT row_to_json($0) as row FROM $1 as tbl WHERE id=$2 AND age>$3 LIMIT $4"
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
	expected := "SELECT row_to_json($0) as row FROM $1 as tbl WHERE id=$2"
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

func TestQuery_ToQuerySelectTwoFields(t *testing.T) {
	query := Query{
		Entity: "users",
		Select: []string{
			"id",
			"name",
		},
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "100",
			},
		},
	}
	expected := "SELECT (select row_to_json(_) as row from (select tbl.$0, tbl.$1) as _) as schemaname FROM $1 as tbl WHERE id=$3"
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
