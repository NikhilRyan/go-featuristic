package db

import (
	"log"
	"sync"
	"time"

	"github.com/nikhilryan/go-featuristic/config"
	"gorm.io/driver/mysql"
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
		cfg, _ := config.LoadConfig(".")
		dsn := getDSN(*cfg)
		db, err = gorm.Open(getDialect(cfg.Driver, dsn), &gorm.Config{
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

func getDSN(cfg config.Config) string {
	switch cfg.Driver {
	case "postgres":
		return config.GetMainDSN(cfg)
	case "mysql":
		return config.GetMainDSN(cfg)
	default:
		log.Fatalf("unsupported database driver: %s", cfg.Driver)
		return ""
	}
}

func getDialect(driver string, dsn string) gorm.Dialector {
	switch driver {
	case "postgres":
		return postgres.Open(dsn)
	case "mysql":
		return mysql.Open(dsn)
	default:
		log.Fatalf("unsupported database driver: %s", driver)
		return nil
	}
}
