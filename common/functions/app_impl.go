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
	err := os.MkdirAll(absPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func (a *AppPath) GenerateAppCode(code, path, fileName, goVer, pkg, extension string) error {
	template, errParse := parseTemplate(code)
	if errParse != nil {
		return errParse
	}

	var tpl bytes.Buffer
	if errExecTemplate := executeTemplate(template, &tpl, goVer, pkg, extension); errExecTemplate != nil {
		return errExecTemplate
	}

	absFilePath := filepath.Join(a.BasePath, path, fileName)
	return writeTofile(absFilePath, []byte(code), &tpl)
}

func writeTofile(path string, code []byte, tpl *bytes.Buffer) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = tpl.WriteTo(file)
	return err
}

func parseTemplate(code string) (*template.Template, error) {
	return template.New("").Parse(code)
}

func executeTemplate(t *template.Template, tpl *bytes.Buffer, goVersion, pkgName, extension string) error {
	data := struct {
		GoVersion string
		Package   string
		Extension string
	}{
		GoVersion: goVersion,
		Package:   pkgName,
		Extension: extension,
	}

	return t.Execute(tpl, data)
}
