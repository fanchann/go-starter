package helpers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type StarterConfigurationYamlVi struct {
	Version  string `yaml:"version"`
	Package  string `yaml:"package"`
	Database string `yaml:"database"`
}

func NewStarterYamlParser(file string) StarterConfigurationYamlVi {
	dataFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("configuration starter.yaml not found!\n")
	}
	var starter StarterConfigurationYamlVi
	if err := yaml.Unmarshal(dataFile, &starter); err != nil {
		log.Fatalf("Error %v", err)
	}
	return starter

}

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
