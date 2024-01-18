package functions

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

type AppPath struct {
	BasePath string
}

func (a *AppPath) CreateAppPath(path string) error {
	absPath := filepath.Join(a.BasePath, path)
	return os.MkdirAll(absPath, os.ModePerm)
}

func (a *AppPath) GenerateAppCode(code interface{}, path, fileName, goVer, pkg string) error {
	tpl, errParse := parseTemplate(code)
	if errParse != nil {
		return errParse
	}

	var tplBuf bytes.Buffer
	if errExecTemplate := executeTemplate(tpl, &tplBuf, goVer, pkg); errExecTemplate != nil {
		return errExecTemplate
	}

	absFilePath := filepath.Join(a.BasePath, path, fileName)
	return writeToFile(absFilePath, tplBuf.Bytes())
}

func writeToFile(path string, code []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(code)
	return err
}

func parseTemplate(code interface{}) (*template.Template, error) {
	return template.New("").Parse(code.(string))
}

func executeTemplate(t *template.Template, tpl *bytes.Buffer, goVersion, pkgName string) error {
	data := struct {
		GoVersion string
		Package   string
	}{
		GoVersion: goVersion,
		Package:   pkgName,
	}

	return t.Execute(tpl, data)
}
