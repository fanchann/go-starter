package code_test

import (
	"os"
	"path"
	"testing"

	"github.com/fanchann/go-starter/common/code"
)

func TestGenerateMysqlCompose(t *testing.T) {
	savedFile := path.Join("gen", "docker-composeMysql.yaml")
	dataFile := code.GenerateDockerCompose("mysql", "root", "root", "test", "3306")

	os.WriteFile(savedFile, []byte(dataFile), os.ModePerm)

}

func TestGeneratePostGresCompose(t *testing.T) {
	savedFile := path.Join("gen", "docker-composePostgres.yaml")
	dataFile := code.GenerateDockerCompose("postgres", "root", "root", "test", "5432")

	os.WriteFile(savedFile, []byte(dataFile), os.ModePerm)

}
