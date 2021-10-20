// Package config provides tools for reading configuration variables
package config

import (
	"github.com/spf13/viper"
)

// Set is a method of config that set config variables.
// This version use viper strucuture of config
func (c *config) Set(file string) error {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	isSet = true
	return nil
}
