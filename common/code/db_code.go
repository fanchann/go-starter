package code

var DBConfigGo = `package database

import (
	"fmt"
	"log"
	"os"

	"{{.Package}}/common/types"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		dbConnection = fmt.Sprintf(types.POSTGRES_CONFIG, d.Host, d.Username, d.Password, d.DBName, d.Port, d.SSlmode)
		dialect = postgres.Open(dbConnection)
	case "mysql":
		dbConnection = fmt.Sprintf(types.MYSQL_CONFIG, d.Username, d.Password, d.Host, d.Port, d.DBName)
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

var DBHelperCode = `package helpers

import (
	"{{.Package}}/common/config"
	"{{.Package}}/common/database"
)

func InitDatabase() *database.DB {
	return &database.DB{
		Driver:   config.GetString("database.driver"),
		Host:     config.GetString("database.host"),
		Username: config.GetString("database.username"),
		Password: config.GetString("database.password"),
		Port:     config.GetInt("database.port"),
		DBName:   config.GetString("database.name"),
		SSlmode:  config.GetString("database.sslmode"),
	}
}
`

var DSN = `package types

var (
	MYSQL_CONFIG    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	POSTGRES_CONFIG = "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai"
)
`
