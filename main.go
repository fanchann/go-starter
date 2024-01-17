package main

import (
	"fmt"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/helpers"
)

var (
	logo = `
	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	`

	driverDBOptions = []string{"mysql", "postgres", "mongodb"}

	appName  = helpers.PromptNamePackage()
	dbDriver = helpers.PromptArrayString("database driver", driverDBOptions)
	dbHost   = helpers.PromptString("database host [example : localhost]")
	dbUser   = helpers.PromptString("username database")
	dbPass   = helpers.PromptString("password database")
	dbName   = helpers.PromptString("database name")
	dbPort   = helpers.PromptInteger("port database")
)

func main() {
	fmt.Printf("%v\n", logo)
	err := app.GoStarter(appName, "", dbDriver, dbHost, dbPort, dbUser, dbPass, dbName)
	helpers.ErrorWithLog(err)
}
