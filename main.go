package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/helpers"
)

var (
	logo = `
	┌─┐┌─┐   ┌─┐┌┬┐┌─┐┬─┐┌┬┐┌─┐┬─┐
	│ ┬│ │───└─┐ │ ├─┤├┬┘ │ ├┤ ├┬┘
	└─┘└─┘   └─┘ ┴ ┴ ┴┴└─ ┴ └─┘┴└─
	`
	showHelp    = flag.Bool("help", false, "show help message")
	starterFile = flag.String("f", "starter.yaml", "starter configuration file")

	starterDefaultConfiguration = map[string]interface{}{
		"version":  "1",
		"package":  "your-name-package",
		"database": "mysql",
	}
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("%v \n", logo)
	newArgs := os.Args[1:]

	if len(newArgs) > 0 {
		command := newArgs[0]
		switch command {
		case "new":
			log.Print("Success generate starter.yaml \n")
			generateConfiguration()
			return
		case "help":
			printHelpMessage()
			return
		}
	}

	processStarterApp(starterFile)

}

func generateConfiguration() error {
	return helpers.NewYamlWriter(starterDefaultConfiguration, ".", "starter.yaml")
}

func processStarterApp(appConfigFile *string) {
	starterConfiguration := helpers.NewStarterYamlParser(*appConfigFile)
	err := app.GoStarter(app.Options{AppName: starterConfiguration.Package, Driver: starterConfiguration.Database})
	helpers.ErrorWithLog(err)
}

func printHelpMessage() {
	fmt.Println("Usage: go-starter [OPTIONS]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  go-starter new      Generate configuration file")
	fmt.Println("  go-starter          Run the application")
	fmt.Println("  go-starter -f=configuration.yaml [default: starter.yaml] Specify a custom configuration file")
}
