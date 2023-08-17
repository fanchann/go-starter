package code

var MainCode = `package main

import (
	"flag"
	"log"
	"os"

	"{{.Package}}/config"
)

var fileConfiguration *string

/*
you can change the configuration file
*/

func init() {
	fileConfiguration = flag.String("c", "config.{{.Extension}}", "Insert your configuration setting")
	flag.Parse()
}

func main() {
	err := config.NewConfiguration(*fileConfiguration).Initialize()
	if err != nil {
		log.Fatalf("Error while read the configuration (%s)", err)
		os.Exit(1)
	}

	db := config.InitDatabase().DatabaseConnection()

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

var MainCodeEnvConfig = `package main

import (
	"flag"
	"log"
	"os"

	"{{.Package}}/lib"
)

var fileConfiguration *string

/*
you can change the configuration file
*/

func init() {
	fileConfiguration = flag.String("c", ".env", "Insert your configuration setting")
	flag.Parse()
}

func main() {
	dbConnection := lib.DatabaseConnection(*fileConfiguration)

	// ping to database
	responseDB, _ := dbConnection.DB()
	if err := responseDB.Ping(); err != nil {
		log.Fatalln("Failed to connect database [%w]", err.Error())
		os.Exit(1)
	}

	// if database successfully connected
	log.Println("[ Success conected to database ]")
}
`
