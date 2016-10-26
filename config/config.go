package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type Config struct {
	Db            DBConfig      `json:"db"`
	Env           string        `json:"env"`
}

type DBConfig struct {
	Master string `json:"master"`
	Slave  string `json:"slave"`
}

var conf *Config

func Init(path string) *Config {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("load config conf failed:", err)
	}
	conf = &Config{}
	err = json.Unmarshal(buf, conf)
	if err != nil {
		log.Fatalln("decode config file failed:", string(buf), err)
	}
	return conf
}

func Instance() *Config {
	if conf == nil {
		Init("app.conf")
	}
	return conf
}

