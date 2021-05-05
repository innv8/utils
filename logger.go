package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iancoleman/orderedmap"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

// InitLogger initializes the logger.
// It requires logFolder and env.
// logFolder is the path of the log directory ending with a slash
// If env is set to dev, the log output will be the stderr.
// If env is set to anything else (e.g staging or prod or nothing),
// the output will be a files, rotated every 24 hours
func InitLogger(logFolder, env string) {
	// if ENV== dev, set output to stderr
	if env == "dev" {
		log.SetOutput(os.Stderr)
		return
	}

	writer, err := rotatelogs.New(
		fmt.Sprintf("%sapp-%s.log", logFolder, "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFolder+"link.log"),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	// if an error occured, print and set the output as stderr
	if err != nil {
		fmt.Printf("unable to initialize writer, logging to stderr")
		log.SetOutput(os.Stderr)
		return
	}

	log.SetOutput(writer)
	return
}

// Requires the following in .env
// APP_NAME
func logger(msg, level string, params ...interface{}) {
	message := fmt.Sprintf(msg, params...)
	if os.Getenv("ENV") == "dev" {
		log.Print(message)
		return
	}

	msgMap := orderedmap.New()
	msgMap.Set("level", level)
	msgMap.Set("time", time.Now().Format("2006-06-02 15:04:05.000000"))
	msgMap.Set("time_stamp", time.Now().UnixNano())
	msgMap.Set("app_name", os.Getenv("APP_NAME"))
	msgMap.Set("message", message)
	msgJSON, _ := json.Marshal(msgMap)
	log.Print(string(msgJSON))
}

// LogINFO logs an INFO message.
func LogINFO(msg string, params ...interface{}) {
	logger(msg, "INFO", params...)
}

// LogError logs an ERROR message
func LogError(msg string, params ...interface{}) {
	logger(msg, "ERROR", params...)
}