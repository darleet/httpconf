package httpconf

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents .yaml config file structure
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

// GetConfig loads config from .yaml (.yml) file into structure.
// Name of .yaml config file should be provided in arguments, if
// file neither exists nor is available to be read from - an error
// is returned.
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
