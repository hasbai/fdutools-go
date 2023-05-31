package xk

import (
	"fdutools-go/captcha"
	"fdutools-go/utils"
	"testing"
)

func TestCaptcha(t *testing.T) {
	xk := New(utils.GetUser())
	err := xk.Login()
	if err != nil {
		t.Fatal(err)
	}
	defer xk.Logout()

	err = captcha.Do(xk.C)
	if err != nil {
		t.Fatal(err)
	}
}
