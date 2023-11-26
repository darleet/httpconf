package httpconf

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Write time.Duration `yaml:"write"`
			Read  time.Duration `yaml:"read"`
			Idle  time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

func GetConfig(fileName string) (*Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf Config
	d := yaml.NewDecoder(f)
	if err := d.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
