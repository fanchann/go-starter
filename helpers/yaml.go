package helpers

import (
	"fmt"
	"io/fs"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func NewYamlWriter(code map[string]interface{}, filePath, fileName string) error {
	codeByte, err := yaml.Marshal(code)
	if err != nil {
		return err
	}
	loc := fmt.Sprintf("%s/%s", filePath, fileName)

	if err := ioutil.WriteFile(loc, codeByte, fs.ModePerm|fs.ModeAppend); err != nil {
		return err
	}

	return nil
}
