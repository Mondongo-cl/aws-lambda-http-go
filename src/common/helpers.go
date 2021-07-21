package common

import (
	"log"
	"os"
)

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("[%s]::Hostname can't be get", err.Error())
		hostname = "<<None>>"
	}
	return hostname
}
