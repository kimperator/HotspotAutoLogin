package Hotspots

import (
	"HotspotAutoLogin/Base"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

type HotspotTelekom struct {
	username  string
	password  string
	cookieJar http.CookieJar
	client    Base.HttpClient
}

func NewHotspotTelekom(username string, password string) *HotspotTelekom {
	jar, _ := cookiejar.New(nil)
	return &HotspotTelekom{username, password, jar, &http.Client{
		Jar: jar,
	},
	}
}

func HttpContent(client Base.HttpClient, url string) string {
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
	var body string
	if r.Body != nil {
		body_r, err := ioutil.ReadAll(r.Body)
		if err == nil {
			body = string(body_r[:len(body_r)])
			defer r.Body.Close()
		}
	}
	return body
}

func (h *HotspotTelekom) CanHandle(wifiEnv Base.WifiEnv) bool {
	ret, _ := regexp.MatchString("Telekom(_HDM)?", wifiEnv.SSID)
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
