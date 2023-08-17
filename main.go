package main

import (
	"fmt"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
)

var (
	configFileOptions = []string{"json", "yaml", "toml", "env"}
	driverDBOptions   = []string{"mysql", "postgres"}

	appName   = helpers.PromptNamePackage()
	configFmt = helpers.PromptArrayString("configuration format", configFileOptions)
	dbDriver  = helpers.PromptArrayString("database driver", driverDBOptions)
	dbHost    = helpers.PromptString("database host [example : localhost]")
	dbUser    = helpers.PromptString("username database")
	dbPass    = helpers.PromptString("password database")
	dbName    = helpers.PromptString("database name")
	dbPort    = helpers.PromptInteger("port database")
)

func main() {

	fmt.Printf("%s \n", types.Logo)
	err := app.GoStarter(appName, configFmt, dbDriver, dbHost, dbPort, dbUser, dbPass, dbName)
	helpers.ErrorWithLog(err)
}
