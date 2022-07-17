package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EnvironmentVariable struct {
	DBServer       string
	DBName         string
	DBUsername     string
	DBPassword     string
	DBPort         int
	DBTimezone     string
	ServiceName    string
	ServicePort    int
	ServiceVersion string
	ServiceDebug   bool
}

func GetDBServer() string {
	return viper.GetString("db.server")
}

func GetDBName() string {
	return viper.GetString("db.name")
}

func GetDBUsername() string {
	return viper.GetString("db.username")
}

func GetDBPassword() string {
	return viper.GetString("db.password")
}

func GetDBPort() int {
	return viper.GetInt("db.port")
}

func GetDBTimezone() string {
	return viper.GetString("db.timezone")
}

func GetServiceName() string {
	return viper.GetString("service.name")
}

func GetServicePort() int {
	return viper.GetInt("service.port")
}

func GetServiceVersion() string {
	return viper.GetString("service.version")
}

func GetServiceDebug() bool {
	return viper.GetBool("Service.debug")
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources/")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("fatal error config file:" + err.Error())
	}

	return nil
}

func GetDBConnections() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", GetDBServer(), GetDBUsername(), GetDBPassword(), GetDBName(), GetDBPort(), GetDBTimezone())

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err = errors.New("failed connect to db:" + err.Error())
		return
	}

	return
}
