package fdu

import (
	"fdutools-go/utils"
	"testing"
)

func TestLogin(t *testing.T) {
	fd := New(utils.GetUser())
	err := fd.Login()
	if err != nil {
		t.Error(err)
	}
	err = fd.Logout()
	if err != nil {
		t.Error(err)
	}
}
