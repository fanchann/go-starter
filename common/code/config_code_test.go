package code_test

import (
	"os"
	"testing"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/helpers"
)

var dirTest = "./gen"
var perm = os.FileMode(0755)

func init() {
	err := os.MkdirAll(dirTest, perm)
	if err != nil {
		panic(err)
	}
}

func TestCreateJson(t *testing.T) {
	s := code.WriteAppConfiguration("json")

	err := os.WriteFile(dirTest+"/settings.json", []byte(s), perm)
	helpers.ErrorWithLog(err)
}

func TestCreateToml(t *testing.T) {
	s := code.WriteAppConfiguration("toml")

	err := os.WriteFile(dirTest+"/settings.toml", []byte(s), perm)
	helpers.ErrorWithLog(err)
}

func TestCreateYaml(t *testing.T) {
	s := code.WriteAppConfiguration("yaml")

	err := os.WriteFile(dirTest+"/settings.yaml", []byte(s), perm)
	helpers.ErrorWithLog(err)
}
