package pgsql

import (
	"fmt"
	"github.com/vsixz/prego/cons"
	"github.com/vsixz/prego/orm/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var db *gorm.DB

func DB(dbConf types.DatabaseConfig) *gorm.DB {
	if dbConf == (types.DatabaseConfig{}) {
		panic("Read mysql database config failed.")
	}
	if db != nil {
		return db
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&timeout=60m",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database)

	var logLevel = logger.Error
	if os.Getenv(cons.GoEnvironment) == string(cons.Debugging) {
		logLevel = logger.Info
	}

	pgsqlConf := postgres.Config{
		DSN: conn,
	}

	gormConf := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, _ = gorm.Open(postgres.New(pgsqlConf), &gormConf)

	if db != nil {
		pool, _ := db.DB()
		pool.SetMaxIdleConns(10)
		pool.SetMaxOpenConns(100)
		pool.SetConnMaxLifetime(time.Hour)
	}
	return db
}