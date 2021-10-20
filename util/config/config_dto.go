// Package config provides tools for reading configuration variables
package config

const (
	// cFile definw a default configuration path and file name to be used
	cFile = "./config.yml"
)

var (
	// c store static config structure
	c ConfigInterface = &config{
		Server: serverConfig{Port: 0},
		Db:     dbConfig{Name: "", User: "", Pass: ""},
	}
	isSet = false
)

// ConfigInterface is a interface thar abstract config type that can be mocket
type ConfigInterface interface {
	// Set is a method of config that set config variables.
	Set(string) error
	// IsSet is a method that return if variables of config was established 
	IsSet() bool
	// GetServer is a method that return server configuration data structure
	GetServer() serverConfig
	// GetDB is a method that return database configuration data structure
	GetDB() dbConfig
}

// config is struct type that describe all config body variables
type config struct {
	Server serverConfig `yaml:"server"`
	Db     dbConfig     `yaml:"db"`
}

// serverConfig is struct type thar describe server config variables
type serverConfig struct {
	Port int `yaml:"port"`
}

// dbConfig is struct type thar describe database config variables
type dbConfig struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// IsSet is a method that return if variables of config was established 
// by reading a file or by other type of configuration
func (c *config) IsSet() bool {
	return isSet
}

// GetServer is a method that return server configuration data structure
func (c *config) GetServer() serverConfig {
	return c.Server
}

// GetDB is a method that return database configuration data structure
func (c *config) GetDB() dbConfig {
	return c.Db
}

// init try to set a config bt a default file path
func init() {
	c.Set(cFile)
}

// GetConfig is a function that return configuration a configuration static variable
func GetConfig() ConfigInterface {
	return c
}
