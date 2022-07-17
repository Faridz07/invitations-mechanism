package database

import (
	"errors"
	"fmt"
	"invitations-mechanism/config"
	"invitations-mechanism/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnections() (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.GetDBUsername(), config.GetDBPassword(), config.GetDBServer(), config.GetDBPort(), config.GetDBName())

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		err = errors.New("failed connect to db:" + err.Error())
		return
	}

	err = Migrate(db)
	if err != nil {
		err = errors.New("failed to migrate table:" + err.Error())
		return
	}

	return
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}
