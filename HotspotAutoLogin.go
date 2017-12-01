package main

import (
	"HotspotAutoLogin/Base"
	"HotspotAutoLogin/Hotspots"
	"flag"
	"fmt"
)

func main() {
	Base.LoadCfg()
	actionPtr := flag.String("action", "login", "which action to perform (login|logout|canHandle)")
	providerPtr := flag.String("provider", "Telekom", "Force using provider (currently only supported method)")

	ssidPtr := flag.String("ssid", "", "WiFi SSID")
	bssidPtr := flag.String("bssid", "", "WiFi BSSID")
	flag.Parse()
	wifiEnv := Base.WifiEnv{*ssidPtr, *bssidPtr}
	hotspot := Hotspots.NewHotspot(*providerPtr, Base.Cfg["TELEKOM_USERNAME"], Base.Cfg["TELEKOM_PASSWORD"])
	if *actionPtr == "login" {
		fmt.Println(hotspot.Login())
	} else if *actionPtr == "logout" {
		fmt.Println(hotspot.Logout())
	} else if *actionPtr == "canHandle" {
		fmt.Println(hotspot.CanHandle(wifiEnv))
	}
}
