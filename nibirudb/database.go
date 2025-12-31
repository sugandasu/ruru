package nibirudb

import (
	"context"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	DB(ctx context.Context) *gorm.DB
	Transaction() Transaction
}
type database struct {
	db *gorm.DB
	tx Transaction
}

func NewDatabaseConnection(cfg *Config) Database {
	var dial gorm.Dialector
	switch cfg.Driver {
	case "postgres":
		dial = postgres.Open(cfg.GetDSN())
	case "mysql":
		dial = mysql.Open(cfg.GetDSN())
	default:
		panic("Database drive not found")
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic("Database session failed to initialize: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Database connection failed: %s" + err.Error())
	}

	if cfg.MaxIdleConnections != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	}

	if cfg.MaxOpenConnections != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	}

	maxLifeTime, err := time.ParseDuration(cfg.MaxConnectionLifeTime)
	if err != nil {
		maxLifeTime = time.Minute * 10
	}
	sqlDB.SetConnMaxLifetime(maxLifeTime)

	maxIdle, err := time.ParseDuration(cfg.MaxConnectionIdleTime)
	if err != nil {
		maxIdle = time.Minute * 5
	}

	sqlDB.SetConnMaxIdleTime(maxIdle)

	if cfg.DebugMode {
		db = db.Debug()
	}

	err = db.Use(otelgorm.NewPlugin())
	if err != nil {
		panic("Database opentelemetry is error: " + err.Error())
	}

	return &database{
		db: db,
		tx: NewTransaction(db),
	}
}
