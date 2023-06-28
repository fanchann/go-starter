package app

import (
	"fmt"
	"strings"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/common/functions"
	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
)

func GoStarter(app, config *string) error {
	folderName := strings.ReplaceAll(*app, "/", "-")
	home := functions.AppPath{BasePath: folderName}

	folders := []string{
		"cmd",
		"common/config",
		"common/database",
		"common/helpers",
		"common/middleware",
		"common/types",
		"docs",
		"models",
		"repository",
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
		{Path: "common/config/", FileName: "load.go", Code: code.LoadConfigCode},
		{Path: "common/database/", FileName: "database.go", Code: code.DBConfigGo},
		{Path: "common/helpers/", FileName: "db.go", Code: code.DBHelperCode},
		{Path: "common/helpers/", FileName: "migrate.go", Code: code.MigrateCode},
		{Path: "common/types/", FileName: "dsn.go", Code: code.DSN},
		{Path: "/", FileName: "go.mod", Code: code.GoMod},
		{Path: "/", FileName: fmt.Sprintf("config.%s", *config), Code: code.WriteAppConfiguration(*config)},
	}

	for _, structure := range appStructure {
		if err := home.GenerateAppCode(structure.Code, structure.Path, structure.FileName, helpers.GetGoVersion(), *app, *config); err != nil {
			return err
		}

	}
	return nil

}
