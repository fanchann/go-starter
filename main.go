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
)

func main() {
	fmt.Printf("%v\n", logo)
	err := app.GoStarter(app.Options{AppName: appName, Driver: dbDriver})
	helpers.ErrorWithLog(err)
}
