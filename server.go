package main

import (
	"invitations-mechanism/config"
	"invitations-mechanism/delivery/web"
	"invitations-mechanism/infrastructure/logger"
	"strconv"
)

var (
	ServerName    string
	ServerPort    int
	ServerVersion string
	ServerDebug   bool
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		logger.LogError(config.LoadConfig, err.Error())
	}

	ServerName = config.GetServerName()
	ServerPort = config.GetServerPort()
	ServerVersion = config.GetServerVersion()
	ServerDebug = config.GetServerDebug()
}

func main() {

	db, err := config.GetDBConnections()
	if err != nil {
		logger.LogError(config.GetDBConnections, err.Error())
	}

	router := web.Router(db)
	router.Run(":" + strconv.Itoa(ServerPort))
}
