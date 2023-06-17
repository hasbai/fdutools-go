package fdu

import (
	"errors"
	"fdutools-go/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Fdu struct {
	C         *Client
	User      *utils.User `json:"user"`
	LoginURL  string      `json:"loginURL"`
	LogoutURL string      `json:"logoutURL"`
}

const loginURL = "https://uis.fudan.edu.cn/authserver/login"
const logoutURL = "https://uis.fudan.edu.cn/authserver/logout"

func New(user *utils.User) *Fdu {
	fdu := Fdu{
		C:         NewClient(),
		User:      user,
		LoginURL:  loginURL,
		LogoutURL: logoutURL,
	}
	return &fdu
}

func (f *Fdu) login(uid string, pwd string) error {
	if uid == "" {
		uid = f.User.Uid
	}
	if pwd == "" {
		pwd = f.User.Pwd
	}

	payload := map[string]string{
		"username": uid,
		"password": pwd,
	}
	resp, err := f.C.Get(f.LoginURL)
	if err != nil {
		return FDUErr(ErrNetwork, err.Error())
	}
	//time.Sleep(time.Millisecond * 200)

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return FDUErr(ErrLogin, err.Error())
	}
	html.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		payload[name] = value
	})
	resp, err = f.C.Post(f.LoginURL, payload)
	if err != nil {
		return FDUErr(ErrNetwork, err.Error())
	}
	//time.Sleep(time.Millisecond * 200)

	if strings.Contains(resp.Header.Get("Expires"), "1970") {
		html, err = goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return FDUErr(ErrParseHTML, err.Error())
		}
		txt := html.Find("#msg").Text()
		if strings.Contains(txt, "密码") {
			return FDUErr(ErrInvalidPassword, txt)
		} else if strings.Contains(txt, "验证码") {
			return FDUErr(ErrNeedCaptcha, txt)
		}
		return FDUErr(ErrLogin, html.Text())
	}
	log.Println("login successful")

	if f.LoginURL == loginURL {
		html, err = goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return FDUErr(ErrParseHTML, err.Error())
		}
		f.User.Name = html.Find(".auth_username span span").Text()
	}

	return nil
}

func (f *Fdu) Login(credential []string) (*utils.User, error) {
	var uid, pwd string
	switch len(credential) {
	case 0:
		uid, pwd = f.User.Uid, f.User.Pwd
	case 2:
		uid, pwd = credential[0], credential[1]
	default:
		return nil, errors.New("invalid credential")
	}
	if !isValidCredential(uid, pwd) {
		return nil, errors.New("invalid credential")
	}

	err := f.login(uid, pwd)
	if err == nil && len(credential) == 2 {
		f.User.Uid, f.User.Pwd = uid, pwd
		utils.SaveConfig()
	}
	return f.User, err
}

func (f *Fdu) Logout() {
	_, err := f.C.Get(f.LogoutURL)
	if err != nil {
		log.Panic(err)
	}
	log.Println("logout successful")
}
