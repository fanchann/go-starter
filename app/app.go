package app

import (
	"fmt"
	"strings"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/common/functions"
	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
)

func GoStarter(app *string) error {
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
			fmt.Printf("Failed to create folder '%s': %v\n", folder, err)
		} else {
			fmt.Printf("Folder '%s' created successfully.\n", folder)
		}
	}

	filloC := []types.AppStructure{
		{Path: "cmd/main.go", Code: code.MainCode},
		{Path: "config/load.go", PackageName: *app, Code: code.LoadConfigCode},
		{Path: "database/database.go", PackageName: *app, Code: code.DBConfigGo},
		{Path: "types/dsn.go", PackageName: *app, Code: code.DSN},
		{Path: "go.mod", PackageName: *app, GoVer: helpers.GetGoVersion(), Code: code.GoMod},
		{Path: "go.sum", Code: code.GoSum},
	}

	for _, r := range filloC {
		err := home.GenerateAppCode(r)
		helpers.ErrorWithLog(err)
	}
	return nil
}
