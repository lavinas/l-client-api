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
	// fileP store a default test file path and name created for unit testes
	fileP = "./testy.yml"
)

var (
	// yamv store a variable structure 
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
	c := GetConfig()
	assert.False(t, c.IsSet())
}

func TestSetError(t *testing.T) {
	c := GetConfig()
	err := c.Set("error.yml")
	assert.Error(t, err)
	assert.EqualValues(t, "open error.yml: no such file or directory", err.Error())
	assert.False(t, c.IsSet())
}

func TestSetOk(t *testing.T) {
	c := GetConfig()
	err := c.Set(fileP)
	assert.Nil(t, err)
	assert.True(t, c.IsSet())
	assert.EqualValues(t, c.GetDB(), yamv.Db)
	assert.EqualValues(t, c.GetServer(), yamv.Server)
}

func TestSetPart1(t *testing.T) {
	type configP struct {
		Server serverConfig `yaml:"server"`
	}
	var yamvp = configP{Server: serverConfig{Port: 100}}
	y, _ = yaml.Marshal(&yamvp)
	ioutil.WriteFile(fileP, y, 0644)
	defer os.Remove(fileP)
	c := GetConfig()
	err := c.Set(fileP)
	assert.Nil(t, err)
	err = c.Set(fileP)
	assert.Nil(t, err)
	assert.EqualValues(t, c.GetServer(), yamvp.Server)
}

func TestSetPart2(t *testing.T) {
	type configP struct {
		Db dbConfig `yaml:"db"`
	}
	var yamvp = configP{Db: dbConfig{Name: "127.0.0.1:3000", User: "test", Pass: "test"}}
	y, _ = yaml.Marshal(&yamvp)
	ioutil.WriteFile(fileP, y, 0644)
	defer os.Remove(fileP)
	c := GetConfig()
	err := c.Set(fileP)
	assert.Nil(t, err)
	err = c.Set(fileP)
	assert.Nil(t, err)
	assert.EqualValues(t, c.GetDB(), yamvp.Db)
}
