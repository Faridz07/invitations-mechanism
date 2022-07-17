package logger

import (
	"os"
	"reflect"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}
}

func LogInfo(context interface{}, message string) {
	log.WithFields(log.Fields{
		"context": getFunctionName(context),
	}).Info(message)
}

func LogError(context interface{}, message string) {
	log.WithFields(log.Fields{
		"context": getFunctionName(context),
	}).Error(message)
}

func getFunctionName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
