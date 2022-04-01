package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Http     `yaml:"http"`
	Log      `yaml:"log"`
	DataBase `yaml:"database"`
}

type Http struct {
	Addr             string `yaml:"addr"`
	DefaultSecretKey string `yaml:"default_secretKey"`
	Csvlog           string `yaml:"csvlog"`
}

type Log struct {
	Level     string `yaml:"level"`
	Filename  string `yaml:"filename"`
	MaxSize   int    `yaml:"max_size"`
	MaxAge    int    `yaml:"max_age"`
	MaxBacket int    `yaml:"max_backet"`
}

type DataBase struct {
	Dbtype       string `yaml:"dbtype"`
	Hostname     string `yaml:"hostname"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Base         string `yaml:"base"`
	PingAttempts int    `yaml:"ping_attempts"`
}

var Conf *Config = &Config{}

func Init(filename string) (err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, Conf)

	if err != nil {
		return err
	}
	return

}
