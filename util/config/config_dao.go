// Package config provides tools for reading configuration variables
package config

import (
	"github.com/spf13/viper"
)

func (c *config) Set(file string) error {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	return nil
}
