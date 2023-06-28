package main

import (
	"flag"
	"fmt"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/common/types"
	"github.com/fanchann/go-starter/helpers"
)

var pkg *string
var cfg *string
var help *bool

func init() {
	pkg = flag.String("pkg", "generateFromGoStarter_", "fill your app-name")
	cfg = flag.String("config", "toml", "generate configuration file")
	help = flag.Bool("help", false, "show all command in go-starter")
	flag.Parse()
}

func main() {
	if *help {
		fmt.Printf("%s\n%s \n", types.Logo, types.Help)
		return
	}

	fmt.Printf("%s \n", types.Logo)
	err := app.GoStarter(pkg, cfg)
	helpers.ErrorWithLog(err)
}
