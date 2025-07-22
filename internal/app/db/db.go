package db

import (
	"fmt"
	"log"
	"os"
	model "test-assignment/internal/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	user := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("db_port")

	dsn := fmt.Sprintf(` host=%s user=%s password=%s dbname=%s port=%s`,
		dbHost, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(string(dsn)),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalln("Fail to connect to DB, error: ", err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Subscription{})
	if err != nil {
		log.Fatalln("Fail to migrate User error: ", err)
	}
}
