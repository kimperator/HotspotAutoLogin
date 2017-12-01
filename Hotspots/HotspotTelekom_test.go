package Hotspots

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"
)

type TelekomClient struct {
	loggedIn bool
}

func (c *TelekomClient) Get(url string) (*http.Response, error) {
	ret := &http.Response{}

	var content string
	if !c.loggedIn {
		content = "Telekom please login"
	}
	ret.Body = ioutil.NopCloser(strings.NewReader(content))
	return ret, nil
}

func (c *TelekomClient) Post(stringUrl string, contentType string, reader io.Reader) (*http.Response, error) {
	ret := &http.Response{}
	var content string

	parsedUrl, _ := url.Parse(stringUrl)

	if strings.Contains(parsedUrl.Path, "login") {
		c.loggedIn = true
	} else {
		c.loggedIn = false
	}
	if c.loggedIn {
		content = "online"
	} else {
		content = "offline"
	}
	ret.Body = ioutil.NopCloser(strings.NewReader(content))
	return ret, nil
}

func TestHotspotTelekomOk(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	ht := HotspotTelekom{"username", "password", jar, &TelekomClient{false}}
	if !ht.NeedsLogin() {
		t.Error("no login needed")
	}
	if !ht.Login() {
		t.Error("unable to login")
	}
	if ht.NeedsLogin() {
		t.Error("login after login needed")
	}
	if !ht.Logout() {
		t.Error("unable to logout")
	}
	if !ht.NeedsLogin() {
		t.Error("no login needed")
	}
}
