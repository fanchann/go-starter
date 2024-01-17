package code

var ErrorLoggerCode = `package helpers

import (
	"log"
	"os"
)

func ErrorLogger(err error) {
	if err != nil {
		log.Fatalf("❌ error :%v \n", err)
		os.Exit(1)
	}
}
`
