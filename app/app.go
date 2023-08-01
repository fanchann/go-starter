package app

import (
	"fmt"
	"strings"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/common/functions"
	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
)

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
		"lib",
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

	appStructure := []types.AppStructure{
		{Path: "cmd/", FileName: "main.go", Code: code.MainCode},
		{Path: "config/", FileName: fmt.Sprintf("%s_reader.go", extension), Code: code.LoadConfigCode},
		{Path: "lib/", FileName: fmt.Sprintf("%s.go", driver), Code: code.DBLib},
		{Path: "/", FileName: "go.mod", Code: code.GoMod},
		{Path: "/", FileName: fmt.Sprintf("config.%s", extension), Code: code.WriteAppConfiguration(extension, host, driver, username, password, dbname, port)},
	}

	for _, structure := range appStructure {
		if err := home.GenerateAppCode(structure.Code, structure.Path, structure.FileName, helpers.GetGoVersion(), app, extension); err != nil {
			return err
		}

	}
	return nil

}
