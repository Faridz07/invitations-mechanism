package logger

import (
	"invitations-mechanism/model"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
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

func LogInfoWithRequest(context interface{}, message string) {
	log.WithFields(log.Fields{
		"request":    context,
		"response":   "",
		"statusCode": "",
	}).Info(message)
}

func LogInfoWithResponse(context *gin.Context, message string) {
	val, _ := context.Get("logger")
	response, _ := context.Get("response")

	logger := val.(model.Logger)

	log.WithFields(log.Fields{
		"request":    logger,
		"response":   response,
		"statusCode": context.Writer.Status(),
	}).Info(message)
}

func LogErrorWithRequest(context *gin.Context, message string) {
	val, _ := context.Get("logger")
	response, _ := context.Get("response")

	logger := val.(model.Logger)
	log.WithFields(log.Fields{
		"request":    logger,
		"response":   response,
		"statusCode": context.Writer.Status(),
	}).Error(message)
}

func LogErrorWithResponse(context *gin.Context, message string) {
	val, _ := context.Get("logger")
	response, _ := context.Get("response")

	logger := val.(model.Logger)
	log.WithFields(log.Fields{
		"request":    logger,
		"response":   response,
		"statusCode": context.Writer.Status(),
	}).Error(message)
}

func LogError(context interface{}, message string) {
	log.WithFields(log.Fields{
		"context": getFunctionName(context),
	}).Error(message)
}

func LogFatal(context interface{}, message string) {
	log.WithFields(log.Fields{
		"context": getFunctionName(context),
	}).Fatal(message)
}

func getFunctionName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
