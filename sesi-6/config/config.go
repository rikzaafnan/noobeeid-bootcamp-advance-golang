package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"database"`
}

type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {

	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	cfg := Config{}

	// process umarshal to struct config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg

	return

}
