package db

import (
	"fmt"
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
)

var (
	database    *gorm.DB
	once        sync.Once
	initialized bool
)

func InitDB() {
	conf := config.GetConf()
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.DatabaseHost,
		conf.DatabasePort,
		conf.DatabaseUsername,
		conf.DatabaseName,
		conf.DatabasePassword,
		conf.DatabaseSSL,
	))

	if err != nil {
		log.Fatalln("Error connecting to DB: ", err)
	}

	db.AutoMigrate(
		&models.User{},
	)

	database = db
	initialized = true
}

func GetDB() *gorm.DB {
	once.Do(func() {
		if initialized == false {
			InitDB()
			initialized = true
		}
	})
	return database
}
