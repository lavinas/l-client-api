// Package config provides tools for reading configuration variables
package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const (
	fileP = "./testy.yml"
)

var (
	yamv = config{
		Server: serverConfig{Port: 8000},
		Db:     dbConfig{Name: "127.0.0.1:3000", User: "test", Pass: "test"},
	}
	y = []byte{}
)

func testMainWrapper(m *testing.M) int {
	y, _ = yaml.Marshal(&yamv)
	err := ioutil.WriteFile(fileP, y, 0644)
	if err != nil {
		panic(err)

	}
	defer os.Remove(fileP)
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func TestConst(t *testing.T) {
	assert.EqualValues(t, cFile, "./config.yml")
}

func TestInit(t *testing.T) {
	c = GetConfig()
	assert.False(t, c.IsSet())
}

func TestSetError(t *testing.T) {
	c = GetConfig()
	err := c.Set("error.yml")
	assert.Error(t, err)
	assert.EqualValues(t, "open error.yml: no such file or directory", err.Error())
	assert.False(t, c.IsSet())
}

func TestIsSet(t *testing.T) {
	c = GetConfig()
	err := c.Set(fileP)
	assert.Nil(t, err)
	assert.True(t, c.IsSet())
}
