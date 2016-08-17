package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigManager struct {
	file string
}

func NewConfig(file string) *ConfigManager {
	return &ConfigManager{file}
}

func (cm *ConfigManager) Load() (*Config, error) {
	data, err := ioutil.ReadFile(cm.file)
	if err != nil {
		return nil, err
	}

	var config *Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

type Config struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}

type Server struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Datasource struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
