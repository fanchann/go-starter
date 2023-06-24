package main

import (
	"flag"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/helpers"
)

var pkg *string
var cfg *string

func init() {
	pkg = flag.String("pkg", "generateFromGoStarter_", "fill your app-name")
	cfg = flag.String("config", "toml", "generate configuration file")
	flag.Parse()
}

func main() {
	err := app.GoStarter(pkg, cfg)
	helpers.ErrorWithLog(err)
}
