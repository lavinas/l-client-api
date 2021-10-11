// Package config provides tools for reading configuration variables
// config uses go.uber.org/config as support main tool
package config

import (
	"os"

	uberconf "go.uber.org/config"
)

const (
	envFile = "CONFIG_PATH"
	cFile   = "./config.yaml"
)

var (
	c config
)

type configInterface interface {
	Get(string) string
}

type config struct {
	c *uberconf.YAML
}

func (c config) Get(s string) string {
	v := c.c.Get(s)
	return v.Source()
}

func init() {
	var err error
	f, err := os.Open(getFile())
	if err != nil {
		panic(err)
	}
	c.c, err = uberconf.NewYAML(uberconf.Source(f))
	if err != nil {
		panic(err)
	}
}

func getFile() string {
	p := os.Getenv(envFile)
	if p == "" {
		return cFile
	}
	return p
}

func GetConfig() configInterface {
	return c
}

func Get(s string) string {
	v := c.c.Get(s)
	return v.Source()
}
