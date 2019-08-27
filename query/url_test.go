package query

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestBreakUrlEmpty(t *testing.T) {
	url := "users"
	all := splitFullURL(url)
	if len(all) != 1 {
		t.Error("Expected 1 not", len(all))
		return
	}

}

func TestBreakUrlSmall(t *testing.T) {
	url := "users?id=eq.1"
	all := splitFullURL(url)
	if len(all) != 2 {
		t.Error("Expected 2 not", len(all))
		printArr(all)
		return
	}

}

func TestBreakUrlThree(t *testing.T) {
	url := "users?id=eq.1&name=eq.Justin&age=gte.100"
	all := splitFullURL(url)
	if len(all) != 4 {
		t.Error("Expected 4 not", len(all))
		return
	}

}

func TestBreakUrlExplicitly(t *testing.T) {
	url := "users?id=eq.1&name=eq.Justin&age=gte.100"
	all := splitFullURL(url)
	if all[0] != "users" {
		t.Error("Unexpected", all[0])
		printArr(all)
		return
	}
	if all[1] != "id=eq.1" {
		t.Error("Unexpected", all[1])
		printArr(all)
		return
	}
	if all[2] != "name=eq.Justin" {
		t.Error("Unexpected", all[2])
		printArr(all)
		return
	}
	if all[3] != "age=gte.100" {
		t.Error("Unexpected", all[3])
		printArr(all)
		return
	}

}

func printArr(arr []string) {
	for _, b := range arr {
		logrus.Errorln(b)
	}
}
