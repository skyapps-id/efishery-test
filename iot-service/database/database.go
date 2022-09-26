package database

import (
	"fmt"
	"iot-service/entity"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort int
	DbName string
}

func Database() (*gorm.DB, error) {

	config := &Config{
		DbUser: "user",
		DbPass: "pass",
		DbHost: "localhost",
		DbPort: 3306,
		DbName: "iot-service",
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&entity.FeedLogs{}); err != nil {
		log.Println(err)
	}

	return db, err

}
