package config

import (
	"gopkg.in/ini.v1"
)

var Config *ini.File

func InitConfig()  {
	configPATH := "./.env"
	cfg , err := ini.Load(configPATH)
	if err != nil {
		panic(err)
	}
	Config = cfg
}
