package config

import (
	"github.com/go-ini/ini"
)

var config *ini.File

func LoadConfig() {
	cfg, err := ini.Load("watson.ini")
	if err != nil {
		panic(err)
	}
	config = cfg
}

func GetBackendDriver() string {
	section, err := config.GetSection("backend")
	if err != nil {
		panic(err)
	}

	if section.HasKey("driver") {
		return section.Key("driver").String()
	} else {
		return "default"
	}
}