package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("config/config.json", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
