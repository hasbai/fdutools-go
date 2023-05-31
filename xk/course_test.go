package xk

import (
	"fdutools-go/utils"
	"testing"
)

func TestQueryCourses(t *testing.T) {
	xk := New(utils.GetUser())
	err := xk.Login()
	if err != nil {
		t.Fatal(err)
	}
	defer xk.Logout()
	courses, err := xk.QueryCourses(Query{
		Name: "微生物与人类健康",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, course := range courses {
		t.Log(course)
	}
}

func TestSelect(t *testing.T) {
	xk := New(utils.GetUser())
	err := xk.Login()
	if err != nil {
		t.Fatal(err)
	}
	defer xk.Logout()
	err = xk.Select(Query{
		Name: "微生物与人类健康",
	})
	if err != nil {
		t.Fatal(err)
	}
}
