package main

import (
	"fmt"
	"invitations-mechanism/config"
	"invitations-mechanism/delivery/web"
	"invitations-mechanism/infrastructure/logger"
	"strconv"
)

var (
	ServiceName    string
	ServicePort    int
	ServiceVersion string
	ServiceDebug   bool
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		logger.LogError(config.LoadConfig, err.Error())
	}

	ServiceName = config.GetServiceName()
	ServicePort = config.GetServicePort()
	ServiceVersion = config.GetServiceVersion()
	ServiceDebug = config.GetServiceDebug()
}

func main() {

	db, err := config.GetDBConnections()
	if err != nil {
		logger.LogError(config.GetDBConnections, err.Error())
	}

	web := web.Router(db)
	logger.LogInfo(main, fmt.Sprintf("%s running at port %d", ServiceName, ServicePort))
	web.Run(":" + strconv.Itoa(ServicePort))
}
