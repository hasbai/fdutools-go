package fdu

import (
	"errors"
	"fdutools-go/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Fdu struct {
	C         *Client
	User      *utils.User
	LoginURL  string
	LogoutURL string
}

func New(user *utils.User) *Fdu {
	jar, _ := cookiejar.New(nil)
	fdu := Fdu{
		C: &Client{http.Client{
			Timeout: time.Second * 10,
			Jar:     jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}},
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

	if resp.StatusCode == 302 {
		log.Println("login successful")
		return nil
	} else {
		return errors.New("login fail\n" + utils.ReadBody(resp.Body))
	}
}

func (f *Fdu) Logout() {
	resp, err := f.C.Get(f.LogoutURL)
	if err != nil {
		log.Panic(err)
	}
	if resp.StatusCode == 302 {
		log.Println("logout successful")
	} else {
		log.Panicf("logout fail\n" + utils.ReadBody(resp.Body))
	}
}
