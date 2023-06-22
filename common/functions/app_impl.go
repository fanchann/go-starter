package functions

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
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
	data := struct {
		GoVersion   string
		PackageName string
	}{
		GoVersion:   format.GoVer,
		PackageName: format.PackageName,
	}

	codeTemplate, err := template.New("").Parse(format.Code)
	if err != nil {
		helpers.ErrorWithLog(err)
	}

	var tpl bytes.Buffer
	if err := codeTemplate.Execute(&tpl, data); err != nil {
		helpers.ErrorWithLog(err)
	}

	absFilePath := filepath.Join(a.BasePath, format.Path)
	errWrite := ioutil.WriteFile(absFilePath, []byte(format.Code), 0644)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
