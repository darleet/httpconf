package httpconf

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
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

func GetConfig(fileName string) (*ServerConfig, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf ServerConfig
	d := yaml.NewDecoder(f)
	if err := d.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
