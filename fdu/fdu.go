package fdu

import (
	"errors"
	"fdutools-go/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)

type Fdu struct {
	C         *Client
	User      *utils.User
	LoginURL  string
	LogoutURL string
}

func New(user *utils.User) *Fdu {
	fdu := Fdu{
		C:         NewClient(),
		User:      user,
		LoginURL:  "https://uis.fudan.edu.cn/authserver/login",
		LogoutURL: "https://uis.fudan.edu.cn/authserver/logout",
	}
	return &fdu
}

func (f *Fdu) Login() error {
	payload := map[string]string{
		"username": f.User.Uid,
		"password": f.User.Pwd,
	}
	resp, err := f.C.Get(f.LoginURL)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 200)

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	html.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		payload[name] = value
	})
	resp, err = f.C.Post(f.LoginURL, payload)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 200)

	if strings.Contains(resp.Header.Get("Expires"), "1970") {
		txt := utils.ReadBody(resp.Body)
		return errors.New("login fail\n" + txt)
	}
	log.Println("login successful")
	return nil
}

func (f *Fdu) Logout() {
	_, err := f.C.Get(f.LogoutURL)
	if err != nil {
		log.Panic(err)
	}
	log.Println("logout successful")
}
