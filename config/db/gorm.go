package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nikhilryan/go-featuristic/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		cfg, err := config.LoadConfig(".")
		if err != nil {
			log.Fatalf("could not load config: %v", err)
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get database instance: %v", err)
		}

		maxOpenConn := 50
		sqlDB.SetMaxOpenConns(maxOpenConn)
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetConnMaxIdleTime(2 * time.Minute)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
	})
	return db
}
