// Package config provides tools for reading configuration variables
// config uses go.uber.org/config as support main tool
package config

import (
	"os"

	uberconf "go.uber.org/config"
)

const (
	cFile = "./config.yaml"
)

var (
	vFile = cFile
	c config
)

type ConfigInterface interface {
	Get(string) string
	SetFile(string) error
}

type config struct {
	c *uberconf.YAML
}

func (c *config) SetFile(file string) error {
	var err error
	vFile = file
	f, err := os.Open(vFile)
	if err != nil {
		return err
	}
	c.c, err = uberconf.NewYAML(uberconf.Source(f))
	if err != nil {
		return err
	}
	return nil
}

func (c config) Get(s string) string {
	if c.c == nil {
		return ""
	}
	v := c.c.Get(s)
	return v.Source()
}

func init() {
	err := c.SetFile(vFile)
	if err != nil {
		c.c = nil
	}
}

func GetConfig() ConfigInterface {
	return &c
}

func Get(s string) string {
	return c.Get(s)
}
