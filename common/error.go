package common

import (
	"errors"
	"log"
)

var (
	RecordNotFound = errors.New("record not found")
)

// AppRecover use for any goroutines catch recover. stop crash app.
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error: ", err)
	}
}
