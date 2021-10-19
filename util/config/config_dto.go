// Package config provides tools for reading configuration variables
package config

const (
	cFile = "./config.yml"
)

var (
	c ConfigInterface = &config{
		Server: serverConfig{Port: 0},
		Db:     dbConfig{Name: "", User: "", Pass: ""},
	}
	isSet = false
)

type ConfigInterface interface {
	Set(string) error
	IsSet() bool
	GetServer() serverConfig
	GetDB() dbConfig
}

type config struct {
	Server serverConfig `yaml:"server"`
	Db     dbConfig     `yaml:"db"`
}

type serverConfig struct {
	Port int `yaml:"port"`
}

type dbConfig struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

func (c *config) IsSet() bool {
	return isSet
}

func (c *config) GetServer() serverConfig {
	return c.Server
}

func (c *config) GetDB() dbConfig {
	return c.Db
}

func init() {
	c.Set(cFile)
}

func GetConfig() ConfigInterface {
	return c
}
