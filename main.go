package main

import (
	"fmt"
)

func main() {
	loadCfg()
	he := HotspotEnv{"Telekom"}
	hst := NewHotspotTelekom(cfg["TELEKOM_USERNAME"], cfg["TELEKOM_PASSWORD"])
	fmt.Println(hst.CanHandle(he))
	fmt.Println(hst.Login())
}
