package tpl

import (
	"core/config"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetTplContent(tplName string) string {
	var tplDir = config.Config().Global.TplDir
	var tplExt = config.Config().Global.TplExt

	tplName, err := filepath.Abs(tplDir + tplName + tplExt)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(tplName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
