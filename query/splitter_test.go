package query

import "testing"

func TestBuildQueryFromURLForSimpleQuery(t *testing.T) {
	url := "/people?age=eq.25"
	query := BuildQueryFromURL(url)
	if query.Entity != "people" {
		t.Log("Got", query.Entity, "and not", "people")
		t.Fail()
	}
	if query.Field != "age" {
		t.Log("Got", query.Field, "and not", "age")
		t.Fail()
	}
	if query.Comparator != "eq" {
		t.Log("Got", query.Comparator, "and not", "eq")
		t.Fail()
	}
	if query.Value != "25" {
		t.Log("Got", query.Value, "and not", "25")
		t.Fail()
	}

}
