// Logging information
package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "main.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Print("Everything is fine!")
	}
}
