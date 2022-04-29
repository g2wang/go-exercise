package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	RestURL string `yaml:"RestURL" env:"REST_URL" evn-default:"http://localhost:8080"`
}

var Cfg Config

func init() {
	err := cleanenv.ReadConfig("config.yml", &Cfg)
	if err != nil {
		panic("bad config")
	}
}
