package db_driver

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type GormDriver struct {
	*gorm.DB
}

func NewGormDriver(dsn string, poolingConfig *SQLPoolingConfig, gormConfig *gorm.Config) (*GormDriver, error) {
	if gormConfig == nil {
		gormConfig = &gorm.Config{TranslateError: true}
	}
	if poolingConfig == nil {
		poolingConfig = new(SQLPoolingConfig).DefaultConfig()
	}

	conn, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDb, err := conn.DB()

	if err = sqlDb.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err != nil {
		return nil, err
	}
	configureSql(sqlDb, poolingConfig)
	return &GormDriver{conn}, nil
}

type SQLPoolingConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (s SQLPoolingConfig) DefaultConfig() *SQLPoolingConfig {
	return &SQLPoolingConfig{
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 5 * time.Minute,
	}
}

func configureSql(sql *sql.DB, poolingConfig *SQLPoolingConfig) {
	sql.SetMaxIdleConns(poolingConfig.MaxIdleConns)
	sql.SetMaxOpenConns(poolingConfig.MaxOpenConns)
	sql.SetConnMaxLifetime(poolingConfig.ConnMaxLifetime)
}
