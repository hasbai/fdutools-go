package grade

import (
	"fdutools-go/utils"
	"fmt"
	"testing"
)

func TestGetGrades(t *testing.T) {
	g := New(utils.GetUser())
	err := g.Login()
	if err != nil {
		t.Fatal(err)
	}
	defer g.Logout()
	err = g.getGrades()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g.grades)
}
