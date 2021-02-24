package orm

import (
	"github.com/vsixz/prego/orm/mysql"
	"github.com/vsixz/prego/orm/pgsql"
	"github.com/vsixz/prego/orm/sqlserver"
	"github.com/vsixz/prego/orm/types"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(conf types.DatabaseConfig) *Database {
	var db *gorm.DB
	switch conf.Driver {
	case types.Mysql:
		db = mysql.DB(conf)
	case types.Sqlserver:
		db = sqlserver.DB(conf)
	case types.Postgresql:
		db = pgsql.DB(conf)
	case types.Sqlite:
		panic("unimplemented.")
	default:
		panic("unsupported database.")
	}
	return &Database{db}
}
