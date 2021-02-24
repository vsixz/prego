package mysql

import (
	"fmt"
	"github.com/vsixz/prego/cons"
	"github.com/vsixz/prego/orm/types"
	"gorm.io/driver/mysql"
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

	mysqlConf := mysql.Config{
		DSN:                      conn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
	}

	gormConf := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, _ = gorm.Open(mysql.New(mysqlConf), &gormConf)

	if db != nil {
		pool, _ := db.DB()
		pool.SetMaxIdleConns(10)
		pool.SetMaxOpenConns(100)
		pool.SetConnMaxLifetime(time.Hour)
	}
	return db
}
