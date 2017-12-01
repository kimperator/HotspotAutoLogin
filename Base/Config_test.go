package Base

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	fmt.Println(os.Getwd())
	Cfg = loadCfg("../HotspotAutoLogin.json.example")
	if Cfg["TELEKOM_USERNAME"] != "myuser" {
		t.Error("unable to parse telekom username")
	}
	if Cfg["TELEKOM_PASSWORD"] != "mypassword" {
		t.Error("unable to parse telekom password")
	}

}
