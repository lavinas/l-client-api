// Package config provides tools for reading configuration variables
// config uses go.uber.org/config as support main tool
package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type varYam1 struct {
	Foo string `yaml:"foo"`
	Foos []int `yaml:"foos"`
	Baz string `yaml:"baz"`
	Bazes int `yaml:"bazes"`
}

type varYam2 struct {
	Foo string `yaml:"foo"`
	Foos []int `yaml:"foos"`
}

type baseYaml struct {
	Key1 varYam1 `yaml:"key1"`
	Key2 varYam2 `yaml:"key2"`
}

const (
	fileP = "./testy.yaml"
)

var (
	yamv = baseYaml{
		Key1: varYam1{
			Foo: "bar", 
			Foos: []int{1, 2},
			Baz: "fal",
			Bazes: 2,
		},
		Key2: varYam2{
			Foo: "bar", 
			Foos: []int{1, 2},
		},
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
	assert.EqualValues(t, cFile, "./config.yaml")
}

func TestInitNil(t *testing.T) {
	assert.Nil(t, c.c)
}

func TestSetFileNil(t *testing.T) {
	c := GetConfig()
	err := c.SetFile("./xxx.yaml")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "open ./xxx.yaml: no such file or directory")
}

func TestGet1(t *testing.T) {
	c := GetConfig()
	err := c.SetFile(fileP)
	assert.Nil(t, err)
	// ret := c.Get("key1.foo")
	// assert.EqualValues(t, yamv.Key1.Foo, ret)
}