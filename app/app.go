package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/common/functions"
	"github.com/fanchann/go-starter/helpers"
)

type AppStructure struct {
	Path     string
	Code     interface{}
	FileName string
}

func GoStarter(app, extension, driver, host string, port int, username, password, dbname string) error {
	folderName := strings.ReplaceAll(app, "/", "-")
	home := functions.AppPath{BasePath: folderName}

	folders := []string{
		"cmd",
		"app/repositories",
		"app/controllers",
		"app/domain/models",
		"app/domain/types",
		"app/routers",
		"app/middlewares",
		"app/services",
		"utils",
		"config",
	}

	for _, folder := range folders {
		err := home.CreateAppPath(folder)
		if err != nil {
			return fmt.Errorf("failed to create layer '%s': [%v]", folder, err)
		} else {
			fmt.Printf("layer '%s' created successfully.\n", folder)
		}
	}

	appStructure := []AppStructure{
		{Path: "cmd/", FileName: "main.go", Code: writeFileCodeSelection(extension, code.MainCodeEnvConfig, code.MainCode)},
		{Path: "config/", FileName: fmt.Sprintf("%s.go", driver), Code: writeFileCodeSelection(extension, code.DBConfigWithEnvSetting, code.DBLib)},
		{Path: pathSelection(extension), FileName: namingFileSelection(extension), Code: writeFileCodeSelection(extension, code.WriteAppConfiguration("env", host, driver, username, password, dbname, port), code.LoadConfigCode)},
		{Path: "/", FileName: "go.mod", Code: writeFileCodeSelection(extension, code.GoModEnv, code.GoMod)},
		{Path: "/", FileName: "docker-compose.yaml", Code: code.GenerateDockerCompose(driver, username, password, dbname, strconv.Itoa(port))},
	}

	if extension == "env" {
		appStructure = append(appStructure, AppStructure{Path: "config/", FileName: "config.go", Code: writeFileCodeSelection(extension, code.ConfigCodeEnv, "")})
	} else {
		appStructure = append(appStructure, AppStructure{
			Path:     "/",
			FileName: fmt.Sprintf("config.%s", extension),
			Code:     code.WriteAppConfiguration(extension, host, driver, username, password, dbname, port),
		})

	}

	for _, structure := range appStructure {

		if err := home.GenerateAppCode(structure.Code, structure.Path, structure.FileName, helpers.GetGoVersion(), app, extension); err != nil {
			return err
		}

	}
	return nil

}

func writeFileCodeSelection(extension, codeA, codeB string) string {
	if extension == "env" {
		return codeA
	}
	return codeB
}

func namingFileSelection(extension string) string {
	if extension == "env" {
		return ".env"
	}
	return fmt.Sprintf("%s_reader.go", extension)
}

func pathSelection(extension string) string {
	if extension != "env" {
		return "config/"
	}
	return "/"
}
