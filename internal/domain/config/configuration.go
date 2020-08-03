package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config is a structure that holds all the configuration
type Config struct {
	Server Server `json:"server" yaml:"server" toml:"server"`
	Limits Limits `json:"limits" yaml:"limits" toml:"limits"`
	DB     DBConf `json:"db" yaml:"db" toml:"db"`
}

// Server is a structure that holds settings for the server
type Server struct {
	Host  string `json:"host" yaml:"host" toml:"host"`
	Port  string `json:"port" yaml:"port" toml:"port"`
	Level int    `json:"level" yaml:"level" toml:"level"`
}

// Limits is a structure that holds settings for all the limits
type Limits struct {
	Login    int    `json:"login" yaml:"login" toml:"login"`
	Password int    `json:"password" yaml:"password" toml:"password"`
	IP       int    `json:"ip" yaml:"ip" toml:"ip"`
	Expire   string `json:"expire" yaml:"expire" toml:"expire"`
}

// DBConf is a struct that holds settings for the project's DB
type DBConf struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	Password string `json:"password" yaml:"password" toml:"password"`
	Name     string `json:"name" yaml:"name" toml:"name"`
	User     string `json:"user" yaml:"user" toml:"user"`
	SSL      string `json:"ssl" yaml:"ssl" toml:"ssl"`
}

// InitConfig is the main method to initialise Config
func InitConfig(cfgPath string) (*Config, error) {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("init config: %s", err)
	}
	cfg := new(Config)
	return cfg, json.Unmarshal(b, cfg)
}
