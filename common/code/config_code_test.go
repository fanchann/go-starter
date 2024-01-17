package code_test

import (
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/common/code"
)

func TestGenConfigMysql(t *testing.T) {
	m := code.ConfigurationFileGenerate("mysql")

	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"config.mysql.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}

func TestGenConfigMongoDB(t *testing.T) {

	m := code.ConfigurationFileGenerate("mongodb")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"config.mongodb.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}

func TestGenConfigPostgres(t *testing.T) {
	m := code.ConfigurationFileGenerate("postgres")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := ioutil.WriteFile(genFolder+"config.postgres.yaml", out, fs.ModePerm|fs.ModeAppend)
	assert.Nil(t, err2)
}
