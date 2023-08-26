package code

var DBLib = `package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MYSQL_CONFIG    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	POSTGRES_CONFIG = "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai"
)

type DB struct {
	Driver   string ` + "`config:\"driver\"`" + `
	Host     string ` + "`config:\"host\"`" + `
	Username string ` + "`config:\"username\"`" + `
	Password string ` + "`config:\"password\"`" + `
	Port     int    ` + "`config:\"port\"`" + `
	DBName   string ` + "`config:\"name\"`" + `
	SSlmode  string ` + "`config:\"sslmode\"`" + `
}

func (d *DB) DatabaseConnection() *gorm.DB {
	var dialect gorm.Dialector
	var dbConnection string

	switch d.Driver {
	case "postgres":
		dbConnection = fmt.Sprintf(POSTGRES_CONFIG, d.Host, d.Username, d.Password, d.DBName, d.Port, d.SSlmode)
		dialect = postgres.Open(dbConnection)
	case "mysql":
		dbConnection = fmt.Sprintf(MYSQL_CONFIG, d.Username, d.Password, d.Host, d.Port, d.DBName)
		dialect = mysql.Open(dbConnection)
	}

	db, errConnect := gorm.Open(dialect, &gorm.Config{})

	if errConnect != nil {
		log.Fatalf("Error while connect to database :[%s]", errConnect)
		os.Exit(1)
	}

	return db
}
`

var DBConfigWithEnvSetting = `package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MYSQL_CONFIG    = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	POSTGRES_CONFIG = "host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=Asia/Shanghai"
)

func DatabaseConnection(configuration Config) *gorm.DB {
	var dialect gorm.Dialector
	var dbConnection string

	switch os.Getenv("db_driver") {
	case "postgres":
		dbConnection = fmt.Sprintf(POSTGRES_CONFIG, configuration.Get("db_host"), configuration.Get("db_username"), configuration.Get("db_password"), configuration.Get("db_name"), configuration.Get("db_port"), configuration.Get("db_sslmode"))
		dialect = postgres.Open(dbConnection)
	case "mysql":
		dbConnection = fmt.Sprintf(MYSQL_CONFIG, configuration.Get("db_username"), configuration.Get("db_password"), configuration.Get("db_host"), configuration.Get("db_port"), configuration.Get("db_name"))
		dialect = mysql.Open(dbConnection)
	}

	db, errConnect := gorm.Open(dialect, &gorm.Config{})
	if errConnect != nil {
		log.Fatalf("Error while connect to database :[%s]", errConnect)
		os.Exit(1)
	}

	return db

}
`
