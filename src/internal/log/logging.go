package logging

import (
	"fmt"
	"log"
	"os"
)

func Logging(returnedErr error) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("logging: cannot create or open logfile, premission denied: %s\n", err.Error())
		} else {
			fmt.Printf("logging: cannot create or open logfile: %s\n", err.Error())
		}
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.Flags())
	logger.Printf("%s\n", returnedErr.Error())
}
