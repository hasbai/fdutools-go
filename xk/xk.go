package xk

import (
	"fdutools-go/fdu"
	"fdutools-go/utils"
	"github.com/PuerkitoBio/goquery"
	"time"
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
	_, err := xk.Fdu.Login([]string{})
	if err != nil {
		return err
	}

	if xk.User.ProfileID == "" {
		resp, err := xk.C.Get("https://xk.fudan.edu.cn/xk/stdElectCourse!innerIndex.action")
		if err != nil {
			return err
		}
		html, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return err
		}
		s := html.Find("input[name='electionProfile.id']").First()
		value, _ := s.Attr("value")
		xk.User.ProfileID = value
		time.Sleep(time.Millisecond * 200)
	}

	_, err = xk.C.Post(
		"https://xk.fudan.edu.cn/xk/stdElectCourse!defaultPage.action",
		map[string]string{"electionProfile.id": xk.User.ProfileID},
	)
	//if resp.StatusCode != 302 {
	//	return errors.New("login failed")
	//}
	return err
}
