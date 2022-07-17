package main

import (
	"fmt"
	"invitations-mechanism/config"
	"invitations-mechanism/delivery/web"
	"invitations-mechanism/infrastructure/database"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/infrastructure/redis"
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
		logger.LogFatal(config.LoadConfig, err.Error())
	}

	ServiceName = config.GetServiceName()
	ServicePort = config.GetServicePort()
	ServiceVersion = config.GetServiceVersion()
	ServiceDebug = config.GetServiceDebug()
}

func main() {

	db, err := database.GetDBConnections()
	if err != nil {
		logger.LogFatal(database.GetDBConnections, err.Error())
		return
	}

	rdc, err := redis.RedisConnect()
	if err != nil {
		logger.LogFatal(redis.RedisConnect, err.Error())
		return
	}

	web := web.Router(db, rdc)
	logger.LogInfo(main, fmt.Sprintf("%s running at port %d", ServiceName, ServicePort))
	web.Run(":" + strconv.Itoa(ServicePort))
}
