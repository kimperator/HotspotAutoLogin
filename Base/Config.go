package Base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Cfg map[string]string

func getConfigFilePath(filename string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	absFileName := filepath.Join(exPath, filename)
	_, err = os.Stat(absFileName)
	if err != nil {
		exPath, _ = os.Getwd()
		absFileName = filepath.Join(exPath, filename)
	}
	return absFileName
}

func loadCfg(absFileName string) map[string]string {
	ret := make(map[string]string)
	var i_cfg map[string]interface{}

	file, _ := os.Open(absFileName)
	b, _ := ioutil.ReadAll(file)
	file.Close()
	json.Unmarshal(b, &i_cfg)
	for itemKey, itemValue := range i_cfg {
		ret[itemKey] = fmt.Sprint(itemValue)
	}
	return ret
}

func LoadCfg() {
	Cfg = loadCfg(getConfigFilePath("HotspotAutoLogin.json"))
}
