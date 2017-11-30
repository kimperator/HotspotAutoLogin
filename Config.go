package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var cfg map[string]string

func loadCfg() {
	filename := "HotspotAutoLogin.json"
	cfg = make(map[string]string)
	var i_cfg map[string]interface{}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	absFileName := filepath.Join(exPath, filename)
	file, _ := os.Open(absFileName)
	b, _ := ioutil.ReadAll(file)
	file.Close()
	json.Unmarshal(b, &i_cfg)
	for itemKey, itemValue := range i_cfg {
		cfg[itemKey] = fmt.Sprint(itemValue)
	}
}
