package database

import (
	"errors"
	"fmt"
	"invitations-mechanism/config"
	"invitations-mechanism/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnections() (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		config.GetDBServer(),
		config.GetDBUsername(),
		config.GetDBPassword(),
		config.GetDBName(),
		config.GetDBPort(),
		config.GetDBTimezone())

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	return db.AutoMigrate(&model.User{}, &model.Invitation{})
}
