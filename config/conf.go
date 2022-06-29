package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `yaml:"TELEGRAM_TOKEN"`
}

func newConfig() *Config {
	return &Config{
		Token: "",
	}
}

func loadConfig() *Config {
	config := newConfig()
	log.Println("Qweqwe")
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		log.Println(err)
		return nil
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil
	}
	return config
}

func GetToken() string {
	conf := loadConfig()
	return conf.Token
}
