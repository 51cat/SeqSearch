package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.SetFlags(log.Ldate | log.Lmicroseconds)
		log.Fatal(err)
		os.Exit(1)
	}

}