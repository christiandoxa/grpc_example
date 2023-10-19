package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// Connect gets connection of postgresql database
func connect() (db *gorm.DB) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, pass, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func Connect() *gorm.DB {
	once.Do(func() {
		instance = connect()
	})
	return instance
}
