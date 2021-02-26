package orm

import (
	"fmt"
	"github.com/vsixz/prego/cons"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm/logger"
)

func newSqlserver(dbConf DatabaseConfig) *Database {
	if dbConf == (DatabaseConfig{}) {
		panic("Read sqlserver database config failed.")
	}

	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database)

	var logLevel = logger.Error
	if os.Getenv(cons.GoEnvironment) == string(cons.Debugging) {
		logLevel = logger.Info
	}

	db, _ := gorm.Open(sqlserver.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logLevel),
	})

	if db != nil {
		pool, _ := db.DB()
		pool.SetMaxIdleConns(10)
		pool.SetMaxOpenConns(100)
		pool.SetConnMaxLifetime(time.Hour)
	}
	return &Database{db}
}
