package database

import (
	"fmt"
	"log"
	"os"

	"github.com/padinky/imperial-fleet/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	log.Println("DB Conn:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ successfully connected to database!")

	err = DB.AutoMigrate(&model.Spaceship{}, &model.Armament{})
	if err != nil {
		log.Fatal(err)
	}

}
