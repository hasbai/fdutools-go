package fdu

import (
	"fdutools-go/utils"
	"fmt"
	"testing"
)

func TestGetGrades(t *testing.T) {
	g := New(utils.GetUser())
	_, err := g.Login([]string{})
	if err != nil {
		t.Fatal(err)
	}
	defer g.Logout()
	grades, err := g.getGrades()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(grades)
}

func TestFdu_GetRank(t *testing.T) {
	g := New(utils.GetUser())
	_, err := g.Login([]string{})
	if err != nil {
		t.Fatal(err)
	}
	defer g.Logout()
	rank, err := g.GetRank()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rank)
}
