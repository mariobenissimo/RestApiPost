package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func IniziazeLogger() {
	// if a want write log in a file.txt // replace with log in elastic search
	// f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("Failed to create logfile" + "log.txt")
	// 	panic(err)
	// }
	// defer f.Close()

	// impostazioni dei log
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
func WriteLogRequestInfo(method string, path string, message string) {
	log.WithFields(log.Fields{
		"method": method,
		"path":   path,
	}).Info(message)
}
func WriteLogRequesWarnWithError(method string, path string, err string, message string) {
	log.WithFields(log.Fields{
		"method": method,
		"path":   path,
		"err":    err,
	}).Warn(message)
}
func WriteLogError(method string, err string, message string) {
	log.WithFields(log.Fields{
		"method": method,
		"err":    err,
	}).Error(message)
}
func WriteLogInfo(method string, message string, messageInfo string) {
	log.WithFields(log.Fields{
		"method":  method,
		"message": message,
	}).Info(messageInfo)
}
