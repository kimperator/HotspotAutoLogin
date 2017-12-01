package Base

import (
	"path/filepath"
	"testing"
)

func TestConfigParser(t *testing.T) {
	Cfg = loadCfg(getConfigFilePath(filepath.Join("..", "HotspotAutoLogin.json.example")))
	if Cfg["TELEKOM_USERNAME"] != "myuser" {
		t.Error("unable to parse telekom username")
	}
	if Cfg["TELEKOM_PASSWORD"] != "mypassword" {
		t.Error("unable to parse telekom password")
	}

}
