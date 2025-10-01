package model

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}

	DB.AutoMigrate(&Task{})
	return nil
}
