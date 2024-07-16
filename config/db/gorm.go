package db

import (
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
		var err error
		dsn := config.GetMainDSN()
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
		if config.IsProd() {
			maxOpenConn = 800
		}
		sqlDB.SetMaxOpenConns(maxOpenConn)
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetConnMaxIdleTime(2 * time.Minute)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
	})
	return db
}
