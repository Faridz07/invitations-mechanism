package config

import (
	"errors"

	"github.com/spf13/viper"
)

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

func GetJWTSecret() []byte {
	return []byte(viper.GetString("jwt.secret"))
}

func GetJWTExpired() int {
	return viper.GetInt("jwt.expired")
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
