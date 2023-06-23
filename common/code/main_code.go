package code

var MainCode = `package main

import (
	"flag"
	"log"
	"os"
	"%s/common/config"
	"%s/common/helpers"
)

var fileConfiguration *string

/*
you can change the configuration file
*/

func init() {
	fileConfiguration = flag.String("c", "config.toml", "Insert your configuration setting")
	flag.Parse()
}

func main() {
	err := config.NewConfiguration(*fileConfiguration).Initialize()
	if err != nil {
		log.Fatalf("Error while read the configuration (%s)", err)
		os.Exit(1)
	}

	db := helpers.InitDatabase().DatabaseConnection()

	// ping to database
	responseDB, _ := db.DB()
	if err := responseDB.Ping(); err != nil {
		log.Fatalln("Failed to connect database [%w]", err.Error())
		os.Exit(1)
	}

	// if database successfully connected
	log.Println("[ Success conected to database ]")
}
`
