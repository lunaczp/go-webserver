package config

import (
	"gopkg.in/gcfg.v1"
	"sync"
)

type Configure struct {
	Global struct {
		Port   int
		TplDir string
		TplExt string
		Env    string
	}
	DB struct {
		Url  string
		Name string
		User string
		Pass string
	}
}

var conf *Configure
var once sync.Once

func LoadConfig() (*Configure, error) {
	var err error

	once.Do(func() {
		conf = &Configure{}
		configFile := "config/site.ini"
		err = gcfg.ReadFileInto(conf, configFile)
	})

	return conf, err
}

func Config() *Configure {
	return conf
}
