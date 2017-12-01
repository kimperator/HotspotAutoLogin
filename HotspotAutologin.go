package main

import (
	"HotspotAutoLogin/Base"
	"HotspotAutoLogin/Hotspots"
	"fmt"
)

func main() {
	Base.LoadCfg()
	we := Base.WifiEnv{"Telekom", "00:11:22:33:44"}
	hst := Hotspots.NewHotspotTelekom(Base.Cfg["TELEKOM_USERNAME"], Base.Cfg["TELEKOM_PASSWORD"])
	fmt.Println(hst.CanHandle(we))
	fmt.Println(hst.Login())
}
