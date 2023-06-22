package main

import (
	"flag"

	"github.com/fanchann/go-starter/app"
	"github.com/fanchann/go-starter/common/code"
	"github.com/fanchann/go-starter/helpers"
)

var pkg *string
var cfg *string

func init() {
	pkg = flag.String("pkg", "generate-from-go-starter", "")
	cfg = flag.String("config", "toml", "")
	flag.Parse()
}

func main() {
	code.WriteAppConfiguration(*cfg)
	err := app.GoStarter(pkg)
	helpers.ErrorWithLog(err)
}
