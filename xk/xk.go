package xk

import (
	"errors"
	"fdutools-go/fdu"
	"fdutools-go/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type XK struct {
	*fdu.Fdu
}

func New(user *utils.User) *XK {
	fd := fdu.New(user)
	fd.LoginURL = "https://xk.fudan.edu.cn/xk/login.action"
	fd.LogoutURL = "https://xk.fudan.edu.cn/xk/logout.action"
	return &XK{
		fd,
	}
}

func (xk *XK) Login() error {
	payload := map[string]string{
		"username": xk.User.Uid,
		"password": xk.User.Pwd,
	}
	resp, err := xk.C.Post(xk.LoginURL, payload)
	if err != nil {
		return err
	}
	if !strings.Contains(resp.Request.URL.Path, "home") {
		return errors.New("xk login failed")
	}

	resp, err = xk.C.Get("https://xk.fudan.edu.cn/xk/stdElectCourse!innerIndex.action")
	if err != nil {
		return err
	}
	if xk.User.ProfileID == "" {
		html, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return err
		}
		s := html.Find("input[name='electionProfile.id']").First()
		value, _ := s.Attr("value")
		xk.User.ProfileID = value
		utils.SaveConfig()
	}

	resp, err = xk.C.Post(
		"https://xk.fudan.edu.cn/xk/stdElectCourse!defaultPage.action",
		map[string]string{"electionProfile.id": xk.User.ProfileID},
	)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("login failed")
	}

	log.Println("xk login success")
	return nil
}
