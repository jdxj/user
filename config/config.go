package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfg *configuration
)

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	cfg = &configuration{}
	return decoder.Decode(cfg)
}

func DB() db {
	return *cfg.DB
}

func Mode() string {
	return *cfg.Mode
}

func Log() log {
	return *cfg.Log
}

type configuration struct {
	DB   *db     `yaml:"db"`
	Mode *string `yaml:"mode"`
	Log  *log    `yaml:"log"`
}

type db struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
}

type log struct {
	Path string `yaml:"path"`
}
