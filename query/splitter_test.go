package query

import "testing"

func TestBuildQueryFromURLForSimpleQuery(t *testing.T) {
	url := "/people?age=eq.25"
	query := BuildQueryFromURL(url)
	if query.Entity != "people" {
		t.Log("Got", query.Entity, "and not", "people")
		t.Fail()
	}

	if len(query.Comparisons) != 1 {
		t.Error("Expected one query not ", len(query.Comparisons))
		return
	}
	if query.Comparisons[0].Field != "age" {
		t.Error("Expected age not ", query.Comparisons[0].Field)
		return
	}
	if query.Comparisons[0].Comparator != "eq" {
		t.Error("Expected eq not ", query.Comparisons[0].Comparator)
		return
	}
	if query.Comparisons[0].Value != "25" {
		t.Error("Expected 25 not ", query.Comparisons[0].Value)
		return
	}

	//if query.Field != "age" {
	//	t.Log("Got", query.Field, "and not", "age")
	//	t.Fail()
	//}
	//if query.Comparator != "eq" {
	//	t.Log("Got", query.Comparator, "and not", "eq")
	//	t.Fail()
	//}
	//if query.Value != "25" {
	//	t.Log("Got", query.Value, "and not", "25")
	//	t.Fail()
	//}

}

//func TestBuildQueryFromURLForMultiQuery(t *testing.T) {
//	url := "people?age=gte.25&active=is.true"
//	query := BuildQueryFromURL(url)
//	if query.Entity != "people" {
//		t.Log("Got", query.Entity, "and not", "people")
//		t.Fail()
//	}
//	if query.Field != "age" {
//		t.Log("Got", query.Field, "and not", "age")
//		t.Fail()
//	}
//	if query.Comparator != "eq" {
//		t.Log("Got", query.Comparator, "and not", "eq")
//		t.Fail()
//	}
//	if query.Value != "25" {
//		t.Log("Got", query.Value, "and not", "25")
//		t.Fail()
//	}
//
//}
