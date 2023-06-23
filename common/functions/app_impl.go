package functions

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fanchann/go-starter/common/types"
)

type AppPath struct {
	BasePath string
}

func (a *AppPath) CreateAppPath(path string) error {
	err := os.Mkdir(a.BasePath+"/"+path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (a *AppPath) GenerateAppCode(format types.AppStructure) error {

	absFilePath := filepath.Join(a.BasePath, format.Path)
	errWrite := ioutil.WriteFile(absFilePath, []byte(format.Code), 0644)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
