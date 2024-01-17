package code_test

import (
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/common/code"
)

var (
	genFolder = "./gen/"
)

func TestGenMysqlCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("mysql")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"mysql.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}

func TestGenPostgresCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("postgres")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"postgres.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}

func TestGenMongoCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("mongodb")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"mongodb.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}
