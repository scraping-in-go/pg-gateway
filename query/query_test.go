package query

import (
	"fmt"
	"testing"
)

func TestQuery_ToURL(t *testing.T) {
	query := Query{
		Entity:      "users",
		Comparisons: []Comparison{},
	}
	expected := "http://localhost:8080/users"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToURLComparison(t *testing.T) {
	query := Query{
		Entity: "users",
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "2",
			},
		},
	}
	expected := "http://localhost:8080/users?id=eq.2"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToURLComparisons(t *testing.T) {
	query := Query{
		Entity: "users",
		Comparisons: []Comparison{
			{
				Field:      "id",
				Comparator: "eq",
				Value:      "2",
			},
			{
				Field:      "age",
				Comparator: "lt",
				Value:      "10",
			},
		},
	}
	expected := "http://localhost:8080/users?id=eq.2&age=lt.10"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToURLSelect(t *testing.T) {
	query := Query{
		Entity: "users",
		Select: []string{
			"id",
			"name",
		},
		Comparisons: []Comparison{},
	}
	expected := "http://localhost:8080/users?select=id,name"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToURLLimit(t *testing.T) {
	query := Query{
		Entity:      "users",
		Comparisons: []Comparison{},
		Limit:       10,
	}
	expected := "http://localhost:8080/users?limit=10"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToURAll(t *testing.T) {
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
				Value:      "1",
			},
			{
				Field:      "name",
				Comparator: "gte",
				Value:      "Justin",
			},
		},
		Limit: 10,
	}
	expected := "http://localhost:8080/users?select=id,name&id=eq.1&name=gte.Justin&limit=10"
	found := query.ToURL("http://localhost:8080")
	if expected != found {
		t.Error("unexpected url from quest")
		t.Error(found)
		t.Error(expected)
		return
	}
}

func TestQuery_ToQuery(t *testing.T) {
	query := Query{
		Entity:      "users",
		Comparisons: []Comparison{},
	}
	expected := "SELECT row_to_json(tbl) as row FROM users as tbl"
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

	if len(binds) != 0 {
		t.Error("too many bind vars. expected 0, got", len(binds))
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
	expected := "SELECT row_to_json(tbl) as row FROM users as tbl WHERE id=$1 LIMIT $2"
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
	expected := "SELECT row_to_json(tbl) as row FROM users as tbl WHERE id=$1 AND age>=$2 LIMIT $3"
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
	expected := "SELECT row_to_json(tbl) as row FROM users as tbl WHERE id=$1"
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
	expected := "SELECT (select row_to_json(_) as row from (select tbl.id, tbl.name) as _) as schemaname FROM users as tbl WHERE id=$1"
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
