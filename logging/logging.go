package logging

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("INFO: %s", message)
}

func LogError(err error, messages ...string) {
	if len(messages) > 0 {
		log.Printf("ERROR: %s, %s", messages, err.Error())
	} else {
		log.Printf("ERROR: %s", err.Error())
	}
}

func Panic(err error, messages ...string) {
	if len(messages) > 0 {
		log.Panicf("ERROR: %s, %s", messages, err.Error())
	} else {
		log.Panicf("ERROR: %s", err.Error())
	}
}
