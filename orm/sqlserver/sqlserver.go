package sqlserver

import (
	"fmt"
	"github.com/vsixz/prego/cons"
	"github.com/vsixz/prego/orm/types"
	"gorm.io/gorm"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func DB(dbConf types.DatabaseConfig) *gorm.DB {
	if dbConf == (types.DatabaseConfig{}) {
		panic("Read sqlserver database config failed.")
	}
	if db != nil {
		return db
	}
	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database)

	var logLevel = logger.Error
	if os.Getenv(cons.GoEnvironment) == string(cons.Debugging) {
		logLevel = logger.Info
	}

	db, _ = gorm.Open(sqlserver.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if db != nil {
		pool, _ := db.DB()
		pool.SetMaxIdleConns(10)
		pool.SetMaxOpenConns(100)
		pool.SetConnMaxLifetime(time.Hour)
	}
	return db
}
