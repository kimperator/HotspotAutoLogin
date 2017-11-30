package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
}

type HotspotEnv struct {
	SSID string
}

type Hotspot interface {
	Login() bool
	Logout() bool
	CanHandle(HotspotEnv) bool
	NeedsLogin() bool
}

type HotspotTelekom struct {
	username  string
	password  string
	cookieJar http.CookieJar
	client    HttpClient
}

func NewHotspotTelekom(username string, password string) *HotspotTelekom {
	jar, _ := cookiejar.New(nil)
	return &HotspotTelekom{username, password, jar, &http.Client{
		Jar: jar,
	},
	}
}

func HttpContent(client HttpClient, url string) string {
	ret := ""
	resp, err := client.Get(url)
	if err == nil {
		ret = HttpBody(resp)
	} else {
		ret = fmt.Sprintln(err)
	}
	return ret

}

func HttpBody(r *http.Response) string {
	body_r, _ := ioutil.ReadAll(r.Body)
	body := string(body_r[:len(body_r)])
	defer r.Body.Close()
	return body
}

func (h *HotspotTelekom) CanHandle(hotspotEnv HotspotEnv) bool {
	ret, _ := regexp.MatchString("Telekom(_[A-Z]+)?", hotspotEnv.SSID)
	return ret
}

func (h *HotspotTelekom) Login() bool {
	loggedIn := false
	if h.NeedsLogin() {
		fmt.Println("found telekom, try to login!")
		var jsonStruct struct {
			Username   string `json:"username"`
			Password   string `json:"password"`
			RememberMe bool   `json:"rememberMe"`
		}
		jsonStruct.Username = h.username
		jsonStruct.Password = h.password
		jsonStruct.RememberMe = true
		jsonData, _ := json.Marshal(jsonStruct)
		resp, _ := h.client.Post("https://hotspot.t-mobile.net/wlan/rest/login", "application/json", bytes.NewReader(jsonData))
		content := HttpBody(resp)
		loggedIn = strings.Contains(content, "online")
		fmt.Println(content)

	} else {
		loggedIn = true
	}

	return loggedIn
}

func (h *HotspotTelekom) NeedsLogin() bool {
	s := HttpContent(h.client, "http://checkip.dyndns.org")
	return strings.Contains(s, "Telekom")
}

func (h *HotspotTelekom) Logout() bool {
	resp, _ := h.client.Post("https://hotspot.t-mobile.net/wlan/rest/logout", "application/json", strings.NewReader(`{"logout":"doit"}`))
	content := HttpBody(resp)
	return strings.Contains(content, "offline")
}

func main() {
	loadCfg()
	he := HotspotEnv{"Telekom"}
	hst := NewHotspotTelekom(cfg["TELEKOM_USERNAME"], cfg["TELEKOM_PASSWORD"])
	fmt.Println(hst.CanHandle(he))
	fmt.Println(hst.Login())
}
