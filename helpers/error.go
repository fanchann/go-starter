package helpers

import (
	"log"
	"os"
)

func ErrorWithLog(err error) {
	if err != nil {
		log.Fatalf("Error while create configuration [%s]", err.Error())
		os.Exit(1)
	}
}
