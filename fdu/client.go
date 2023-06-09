package fdu

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	http.Client
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	c := &Client{http.Client{
		Timeout: time.Second * 10,
		Jar:     jar,
	}}
	c.AllowRedirect()
	return c
}

const UA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/91.0.4472.114 Safari/537.36"

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", UA)
	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode >= 400 {
		return resp, errors.New("status " + resp.Status)
	}
	return c.repeatLogin(resp)
}

func (c *Client) Get(reqURL string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", reqURL, nil)
	return c.do(req)
}

func (c *Client) Post(reqURL string, body map[string]string) (*http.Response, error) {
	data := url.Values{}
	for k, v := range body {
		data.Set(k, v)
	}
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.do(req)
}

func (c *Client) PostJSON(reqURL string, body map[string]string) (*http.Response, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", reqURL, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

func (c *Client) AllowRedirect() {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return nil
	}
}

func (c *Client) DisallowRedirect() {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}

func (c *Client) repeatLogin(resp *http.Response) (*http.Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	const text = "当前用户存在重复登录的情况，已将之前的登录踢出："
	if doc.Find("h2").First().Text() == text {
		href, _ := doc.Find("a").First().Attr("href")
		return c.Get(href)
	} else {
		resp.Body = io.NopCloser(bytes.NewReader(body))
		return resp, nil
	}
}
