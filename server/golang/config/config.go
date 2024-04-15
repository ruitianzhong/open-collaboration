package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Port        string       `yaml:"port"`
	Db          *Db          `yaml:"db"`
	Auth        *Auth        `yaml:"auth"`
	Translation *Translation `yaml:"translation"`
}

type Translation struct {
	Appid string `yaml:"appid"`
	Key   string `yaml:"key"`
}

type Db struct {
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	Port     string `yaml:"port"`
	Address  string `yaml:"address"`
}

type Auth struct {
	SessionKey string `yaml:"sessionKey"`
}

func ReadConfig() *Config {
	var config Config
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("failed to read config.yaml")

	}
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}
