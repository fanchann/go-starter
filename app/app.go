package app

import (
	"fmt"

	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/common/functions"
	"github.com/fanchann/go-starter/helpers"
)

type Options struct {
	AppName string
	Driver  string
}

type AppStructure struct {
	Path     string
	Code     interface{}
	FileName string
}

func GoStarter(opt Options) error {
	// basePath := strings.ReplaceAll(opt.AppName, "/", "-")
	restDirectoryList := []string{
		"cmd",
		"api",
		"db/migrations",
		"internals/config",
		"internals/delivery/http",
		"internals/delivery/messaging",
		"internals/gateway",
		"internals/models",
		"internals/repository",
		"internals/usecase",
		"internals/helpers",
		"tests",
	}
	return generateStructure(opt, restDirectoryList)
}

func generateStructure(opt Options, DirectoryList []string) error {
	basePath := "./"
	home := functions.AppPath{BasePath: basePath}
	for _, dir := range DirectoryList {
		err := home.CreateAppPath(dir)
		if err != nil {
			return fmt.Errorf("failed to create layer '%s': [%v]", dir, err)
		} else {
			fmt.Printf("layer '%s' created successfully.\n", dir)
		}
	}

	structures := []AppStructure{
		{Path: "cmd/", Code: code.NewDatabaseOptions(opt.Driver).MainApp, FileName: "main.go"},
		{Path: basePath, Code: code.NewDatabaseOptions(opt.Driver).Dependencies, FileName: "go.mod"},
		{Path: "internals/config/", Code: code.NewDatabaseOptions(opt.Driver).DatabaseDriver, FileName: fmt.Sprintf("%s.go", opt.Driver)},
		{Path: "internals/config/", Code: code.ViperConfiguration, FileName: "viper.go"},
		{Path: "internals/helpers/", Code: code.ErrorLoggerCode, FileName: "error.go"},
	}

	// code generator
	for _, structure := range structures {
		if errGenCode := home.GenerateAppCode(structure.Code, structure.Path, structure.FileName, helpers.GetGoVersion(), opt.AppName); errGenCode != nil {
			return errGenCode
		}
	}

	// config generator
	config := []AppStructure{
		{Path: basePath, Code: code.NewDatabaseOptions(opt.Driver).ConfigurationFile, FileName: "config.dev.yaml"},
		{Path: basePath, Code: code.NewDatabaseOptions(opt.Driver).DockerCompose, FileName: "docker-compose.yaml"},
	}

	for _, cfgGen := range config {
		if errGenConfig := helpers.NewYamlWriter(cfgGen.Code.(map[string]interface{}), cfgGen.Path, cfgGen.FileName); errGenConfig != nil {
			return errGenConfig
		}
	}
	return nil
}
