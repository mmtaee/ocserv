package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// DBConfig database configs
type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

var db *gorm.DB

// Connect connecting to database with configs.
func Connect(cfg *DBConfig) {
	var err error
	db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		log.Printf("Database is not ready: %v\n", err)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}
	log.Println("Database connection established")
}

// Connection get database connection
func Connection(debug ...bool) *gorm.DB {
	if len(debug) > 0 && debug[0] {
		db.Debug()
		db.Logger = logger.Default.LogMode(logger.Info)
	}
	return db
}

// Close database connection closing
func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
